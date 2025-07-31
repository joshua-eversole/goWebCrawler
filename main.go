package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load() // Load environment variables from .env file

	portString := os.Getenv(".env")
	if portString == "" {
		log.Fatal("PORT environment variable is not set")
	}

	router := chi.NewRouter()

	// so people can make requests to our server from the router
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of the browsers
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)
	router.Mount("/v1", v1Router)
	// Using the router handler and the port from the environment variable
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	fmt.Println("Server will run on port:", portString)

	// Should skip this if everything runs correctly
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
