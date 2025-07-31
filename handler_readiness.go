package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	// Respond with a simple message indicating the service is ready
	respondWithJSON(w, 200, struct{}{})

}
