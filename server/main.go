package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/yona3/go-auth-sample/router"
)

func main() {
	// load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router.Init()
	http.ListenAndServe(":8080", nil)
}
