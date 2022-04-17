package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/open-wallstreet/go-avanza/avanza"
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
	overview, err := client.Account.GetOverview(context.Background())
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println(overview)

	/* positions, err := client.Account.GetDealsAndOrders(context.Background())
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println(positions) */
	d := models.Date(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))
	transactions, err := client.Account.GetTransactions(context.Background(), &models.GetTransactionsParams{
		AccountID: overview.Accounts[2].AccountId,
		FromDate:  &d,
	})
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println(transactions)
}
