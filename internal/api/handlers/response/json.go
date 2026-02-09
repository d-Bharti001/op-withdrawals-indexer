package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type ResponseEnvelope struct {
	Message string `json:"message"`
	Result  any    `json:"result"`
}

type ErrorEnvelope struct {
	Error string `json:"error"`
}

func SendJSON(w http.ResponseWriter, status int, message string, result any) error {
	w.WriteHeader(status)

	if status == http.StatusNoContent {
		return nil
	}

	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(ResponseEnvelope{
		Message: message,
		Result:  result,
	})
}

func SendErrorResponse(w http.ResponseWriter, status int, message string) error {
	if status == http.StatusInternalServerError {
		log.Println("Internal server error :::", message)
		message = "Internal server error"
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(ErrorEnvelope{
		Error: message,
	})
}
