package main

import (
	"context"
	"github.com/joho/godotenv"
	goavanza "github.com/open-wallstreet/go-avanza/avanza"
	"log"
	"os"
	"time"
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
	client := goavanza.New(goavanza.WithDebug(true))
	_, err := client.Auth.Authenticate(context.Background(), username, password, totpSecret)
	if err != nil {
		log.Fatalf(err.Error())
	}

	timeout, _ := context.WithTimeout(context.Background(), 90*time.Second)
	_, quotes, err := client.Websocket.StreamQuotes(timeout, "19000") // 19000 = USD/SEK
	if err != nil {
		log.Fatalf(err.Error())
	}
	for q := range quotes {
		log.Println(q)
		log.Println(q.Data.Updated.ToTime().String())
		log.Println(q.Data.LastUpdated.ToTime().String())
	}
}
