package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//Response Currently using this for test, will write a more robust system
//for the API
type Response struct {
	Message string `json:"message"`
}

//ErrorWithJSON Responds with An Error Message
func ErrorWithJSON(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	fmt.Fprintf(w, "{ message: %q }", message)
}

//ResponseWithJSON Responds with a Successful JSON
func ResponseWithJSON(w http.ResponseWriter, json []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(json)
}

func NewResponse(message string) ([]byte, error) {
	successResponse := &Response{
		Message: message,
	}
	successResponseJSON, err := json.Marshal(successResponse)
	if err != nil {
		return successResponseJSON, err
	}
	return successResponseJSON, nil
}
