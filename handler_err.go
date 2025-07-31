package main

import "net/http"

func handlerErr(w http.ResponseWriter, r *http.Request) {
	// Respond with a simple message indicating an error occurred
	respondWithError(w, http.StatusInternalServerError, "An unexpected error occurred")
}
