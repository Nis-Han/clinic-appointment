package utils

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Hash(s string) string {
	inputBytes := []byte(s)

	// Create a new SHA-256 hash
	hash := sha256.New()

	// Write the input bytes to the hash
	hash.Write(inputBytes)

	// Get the hash sum as bytes
	hashSum := hash.Sum(nil)

	// Convert the hash sum to a hexadecimal string
	hashString := fmt.Sprintf("%x", hashSum)

	return hashString
}

func GenerateUserKey(first_name string, last_name string, contact_no string, pass string) string {
	return Hash(first_name + last_name + contact_no + Hash(pass))
}

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5XX error", msg)
	}
	type errResponse struct {
		Error string `json:"error"`
	}

	RespondWithJSON(w, code, errResponse{Error: msg})
}

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
