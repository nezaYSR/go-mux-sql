package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/gorilla/mux"
	"github.com/nezaYSR/go-mux-sql/pkg/routes"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	r := mux.NewRouter()
	routes.RegisterScrollStoreRoutes(r)
	http.Handle("/", r)
	url := fmt.Sprintf("localhost:%s", port)
	log.Fatal(http.ListenAndServe(
		url, r,
	))
}
