package handlerError

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func Exec(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	responseError := Response{Message: message}
	if err := json.NewEncoder(w).Encode(responseError); err != nil {
		panic(err)
	}
}
