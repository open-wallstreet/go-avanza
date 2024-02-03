package main

import (
	"context"
	"github.com/joho/godotenv"
	goavanza "github.com/open-wallstreet/go-avanza/avanza"
	"github.com/open-wallstreet/go-avanza/avanza/websocket"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
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
	log, _ := zap.NewDevelopment()
	client := goavanza.New(goavanza.WithDebug(false), goavanza.WithLogger(log.Sugar()))
	_, err := client.Auth.Authenticate(context.Background(), username, password, totpSecret)
	if err != nil {
		log.Fatal(err.Error())
	}

	// timeout, _ := context.WithTimeout(context.Background(), 90*time.Second)
	err = client.Websocket.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)

	client.Websocket.Subscribe(websocket.QuotesSubscriptionPath, "19000")

	for {
		select {
		case <-sigint:
			return
		case <-client.Websocket.Error():
			return
		case out, more := <-client.Websocket.Output():
			if !more {
				return
			}
			switch out.(type) {
			default:
				log.Sugar().Warnf("unknown message type: %T", out)
			}
		}
	}
	/*_, quotes, err := client.Websocket.StreamQuotes(timeout, "19000") // 19000 = USD/SEK
	if err != nil {
		log.Fatalf(err.Error())
	}
	for q := range quotes {
		log.Println(q)
		log.Println(q.Data.Updated.ToTime().String())
		log.Println(q.Data.LastUpdated.ToTime().String())
	}*/
}
