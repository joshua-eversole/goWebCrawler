package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load() // Load environment variables from .env file

	portString := os.Getenv(".env")
	if portString == "" {
		log.Fatal("PORT environment variable is not set")
	}

	router := chi.NewRouter()
	// Using the router handler and the port from the environment variable
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	fmt.Println("Server will run on port:", portString)

	// Should skip this if everything runs correctly
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf(err)
	}

}
