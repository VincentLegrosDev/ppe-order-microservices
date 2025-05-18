package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"ppe4peeps.com/services/order/cmd/server"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("SERVER_PORT")

	server := server.SetupRouter()
	server.Run(fmt.Sprintf(":%s", port))
}
