package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("responding with 5xx error:", msg)
	}

	type errResponse struct {
		Error string `json:"error`
	}

	respondWithJSON(w, code, errResponse{Error: msg})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	response := map[string]string{"error": msg}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println("failed to encode JSON:", err)
	}
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Println("failed to marshal JSON:", err)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}
