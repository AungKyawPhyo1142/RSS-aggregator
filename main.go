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
	fmt.Println("Hello, World!")

	godotenv.Load(".env")

	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("PORT is not set")
	} else {
		fmt.Println("PORT is set to", portString)
	}

	router := chi.NewRouter()

	server := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server is running on port %s", portString)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
