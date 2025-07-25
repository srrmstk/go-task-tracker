package helpers

import (
	"encoding/json"
	"net/http"
)

func JsonError(w http.ResponseWriter, err string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{
		"error": err,
	})
}
