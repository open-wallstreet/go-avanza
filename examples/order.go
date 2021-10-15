package main

import (
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
	goavanza "github.com/open-wallstreet/go-avanza"
	"go.uber.org/zap"
)

func main() {
	godotenv.Load()
	log, _ := zap.NewDevelopment()
	logger := log.Sugar()
	api := goavanza.NewClient(logger)

	err := api.Authenticate()
	if err != nil {
		logger.Panic(err)
	}

	websocketApi := goavanza.NewWebsocket(api, logger, goavanza.NewAvanzaWebsocketOptions())

	go websocketApi.Listen()

	logger.Infof("account id %s", os.Getenv("AVANZA_ACCOUNT_ID"))

	options := &goavanza.OrderOptions{
		AccountId:   os.Getenv("AVANZA_ACCOUNT_ID"),
		OrderbookId: "5452",
		OrderType:   goavanza.ORDER_TYPE_BUY,
		Price:       60,
		ValidUntil:  time.Now(),
		Volume:      1,
	}
	res, err := api.PlaceOrder(
		options,
	)
	if err != nil {
		logger.Panic(err)
	}
	logger.Info(res)

	websocketApi.Subscribe([]string{res.OrderID})

	options.Price = 62
	websocketApi.Unsubscribe([]string{res.OrderID})

	// api.EditOrder(goavanza.STOCK, res.OrderID, options)

	// time.Sleep(5 * time.Second)

	// api.DeleteOrder(os.Getenv("AVANZA_ACCOUNT_ID"), res.OrderID)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	for {
		select {
		case <-interrupt:
			logger.Info("interrupt")
			api.Close()
			return
		}
	}
}
