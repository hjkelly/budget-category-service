package views

import (
	"encoding/json"
	"github.com/hjkelly/budget-category-service/common"
	"io/ioutil"
	"log"
	"net/http"
)

func getRequestBody(r *http.Request, data interface{}) error {
	// Read all data into a byte array.
	bytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	// Try to unmarshal into whatever type we were given.
	err = json.Unmarshal(bytes, data)
	if err != nil {
		return err
	}

	return nil
}

func sendDataAsJSON(w http.ResponseWriter, data interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func sendParseError(w http.ResponseWriter) {
	err := common.APIError{
		Message: "The request body you provided wasn't valid JSON.",
	}
	sendDataAsJSON(w, err, 400)
}

func sendServerError(w http.ResponseWriter, err error, context string) {
	// TODO: This isn't helpful. Write a separate function for logging an error, so this one can truly just send a response.
	log.Printf("ERROR: %s", err.Error())
	apiErr := common.APIError{
		Message: "Something went wrong on our end. Try again later!",
	}
	sendDataAsJSON(w, apiErr, 500)
}
