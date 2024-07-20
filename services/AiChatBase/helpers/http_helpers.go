package helpers

import (
	"encoding/json"
	"net/http"
)

type httpRejection struct {
	Code  int    `json:"code"`
	Cause string `json:"cause"`
}

type booleanResponse struct {
	Response bool `json:"response"`
}

func createRejection(code int, message string) *httpRejection {
	var rejection *httpRejection = new(httpRejection)
	rejection.Code = code
	rejection.Cause = message

	return rejection
}

func WriteRejection(response http.ResponseWriter, code int, message string) {
	var rejection *httpRejection = createRejection(code, message)

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(code)

	json.NewEncoder(response).Encode(rejection)
}

func WriteBooleanResponse(response http.ResponseWriter, value bool) {
	var boolean_response *booleanResponse = new(booleanResponse)
	boolean_response.Response = value

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(200)

	json.NewEncoder(response).Encode(boolean_response)
}
