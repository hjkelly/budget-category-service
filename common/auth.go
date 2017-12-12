package common

import (
	"encoding/json"
	"github.com/auth0-community/go-auth0"
	"gopkg.in/square/go-jose.v2"
	"log"
	"net/http"
)

type JWTValidator struct {
	JWKS_URI           string
	AUTH0_API_ISSUER   string
	AUTH0_API_AUDIENCE []string
}

func (v JWTValidator) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	client := auth0.NewJWKClient(auth0.JWKClientOptions{URI: v.JWKS_URI})
	audience := v.AUTH0_API_AUDIENCE

	configuration := auth0.NewConfiguration(client, audience, v.AUTH0_API_ISSUER, jose.RS256)
	validator := auth0.NewValidator(configuration)

	token, err := validator.ValidateRequest(r)

	if err != nil {
		log.Printf("Missing or invalid token: %s\n", token)
		log.Printf("Technical error: %s\n", err.Error())
		log.Println(err)

		response := APIError{
			Message: "Missing or invalid token.",
		}

		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)

	} else {
		next(w, r)
	}
}
