package main

import (
	"context"
	"github.com/joho/godotenv"
	avanza "github.com/open-wallstreet/go-avanza"
	"github.com/open-wallstreet/go-avanza/avanza/models"
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
	client := avanza.New(avanza.WithDebug(true))
	defer client.Close()
	_, err := client.Auth.Authenticate(context.Background(), username, password, totpSecret)
	if err != nil {
		log.Fatalf(err.Error())
	}

	orderBook, err := client.Order.PlaceOrder(context.Background(), &models.PlaceOrderParams{
		AccountID:   os.Getenv("AVANZA_ACCOUNT_ID"),
		OrderBookID: "5452",
		OrderType:   models.OrderTypeBuy,
		Price:       65,
		ValidUntil:  models.Date(time.Now()),
		Volume:      1,
	})
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println(orderBook)

}
