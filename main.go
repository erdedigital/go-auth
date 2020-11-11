package main

import (
	"log"
	"os"

	"github.com/erdedigital/go-users/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	var port = os.Getenv("PORT")

	if os.Getenv("PORT") == "" {
		port = "8000"
	}
	routes.SetupServer(port)
}
