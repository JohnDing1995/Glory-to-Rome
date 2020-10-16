package utils

import (
	"encoding/json"
	"net/http"
)

func GenerateResponse(status bool, message string, jsonBody interface{}) map[string]interface{} {
	return map[string]interface{}{
		"success": status,
		"message": message,
		"data":    jsonBody,
	}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
