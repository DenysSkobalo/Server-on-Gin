package main

import (
	"api/config"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)


func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %s", err.Error())
	}

	port := os.Getenv("PORT")
	timeout := os.Getenv("TIMEOUT")

	timeoutDuration, err := time.ParseDuration(timeout)
	if err != nil {
		log.Fatalf("Error parsing timeout duration: %s", err.Error())
	}

	cnfg := config.Config{
		Port:    port,
		Timeout: timeoutDuration,
	}

	server := config.NewServer(cnfg)

	if err := server.Run(); err != nil {
		log.Fatalf("Error starting the server: %s", err.Error())
	}
}