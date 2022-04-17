package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/open-wallstreet/go-avanza/avanza"
	"log"
	"os"
)

func main() {
	_ = godotenv.Load()
	totpSecret := os.Getenv("AVANZA_TOTP_SECRET")
	if totpSecret == "" {
		log.Fatalf("AVANZA_TOTP_SECRET environment variable not set")
	}
	username := os.Getenv("AVANZA_USERNAME")
	if username == "" {
		log.Fatalf("AVANZA_USERNAME environment variable not set")
	}
	password := os.Getenv("AVANZA_PASSWORD")
	if password == "" {
		log.Fatalf("AVANZA_PASSWORD environment variable not set")
	}
	client := avanza.New(avanza.WithDebug(true))
	defer client.Close()
	authenticate, err := client.Auth.Authenticate(context.Background(), username, password, totpSecret)
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println(authenticate.CustomerId)
}
