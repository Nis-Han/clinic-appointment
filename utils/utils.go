package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(500)
		log.Printf("Failed to marshal object to JSON with error, %v", err)
	}

	w.Header().Add("Content-Type", "appication/json")
	w.WriteHeader(code)
	w.Write(data)

}
