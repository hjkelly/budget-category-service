package common

import (
	"encoding/json"
	"errors"
	auth0 "github.com/auth0-community/go-auth0"
	jose "gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
	"log"
	"net/http"
	"sync"
)

var instance *Auth
var once sync.Once

func GetAuth() *Auth {
	once.Do(func() {
		instance = &Auth{ // TODO: Pull from env
			JWKS_URI:           "https://zerobalancebudget.auth0.com/.well-known/jwks.json",
			AUTH0_API_ISSUER:   "https://zerobalancebudget.auth0.com/",
			AUTH0_API_AUDIENCE: []string{"https://api.zerobalancebudget.com"},
		}
	})
	return instance
}

// Handle authentication and authorization based on our Auth0 settings.
type Auth struct {
	JWKS_URI           string
	AUTH0_API_ISSUER   string
	AUTH0_API_AUDIENCE []string

	validator *auth0.JWTValidator
}

func (a Auth) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	token, err := a.getToken(r)

	if err != nil {
		a.sendTokenErrorResponse(w, token, err)
	} else {
		next(w, r)
	}
}

// Get the string representing the user from the JWT.
func (a Auth) GetSub(r *http.Request) (string, error) {
	token, err := a.getToken(r)
	if err != nil {
		return "", err
	}

	claims := map[string]interface{}{}
	err = a.validator.Claims(r, token, &claims)
	if err != nil {
		return "", err
	}

	sub, found := claims["sub"]
	if !found {
		return "", errors.New("Token had no 'sub' claim")
	}
	stringSub, ok := sub.(string)
	if !ok {
		return "", errors.New("Token 'sub' claim was somehow not a string.")
	}
	return stringSub, nil
}

func (a Auth) CheckAuthorization(r *http.Request, scopes ...string) error {
	_, err := a.getToken(r)
	if err != nil {
		return err
	}

	// TODO: Check scopes... somehow?
	// claims := map[string]interface{}{}
	// err := validator.Claims(r, token, &claims)
	// if err != nil {
	// 	return err
	// }
	return nil
}

// Make sure the validator is ready.
func (a Auth) prepValidator() {
	client := auth0.NewJWKClient(auth0.JWKClientOptions{URI: a.JWKS_URI})
	audience := a.AUTH0_API_AUDIENCE
	configuration := auth0.NewConfiguration(client, audience, a.AUTH0_API_ISSUER, jose.RS256)
	a.validator = auth0.NewValidator(configuration)
}

func (a Auth) getToken(r *http.Request) (*jwt.JSONWebToken, error) {
	a.prepValidator()
	log.Println(r)
	return a.validator.ValidateRequest(r)
}

func (a Auth) sendTokenErrorResponse(w http.ResponseWriter, token *jwt.JSONWebToken, err error) {
	log.Printf("Missing or invalid token: %#v\n", token)
	log.Printf("Technical error: %s\n", err.Error())
	log.Println(err)

	response := APIError{
		Message: "Missing or invalid token.",
	}

	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(response)
}
