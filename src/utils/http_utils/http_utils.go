package http_utils

import (
	"encoding/json"
	"net/http"

	"github.com/sampado/bookstore_utils-go/rest_errors"
)

func RespondJson(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

func RespondJsonError(w http.ResponseWriter, err rest_errors.RestError) {
	RespondJson(w, err.Status(), err)
}
