package main

import (
	"os"
	"time"

	"github.com/joho/godotenv"
	goavanza "github.com/open-wallstreet/go-avanza"
	"go.uber.org/zap"
)

func main() {
	godotenv.Load()
	log, _ := zap.NewDevelopment()
	logger := log.Sugar()
	api := goavanza.NewApi(logger)

	err := api.Authenticate()
	if err != nil {
		logger.Panic(err)
	}

	overview, err := api.GetOverview()
	if err != nil {
		panic(err)
	}

	logger.Info(overview)

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

	options.Price = 62

	api.EditOrder(goavanza.STOCK, res.OrderID, options)
	api.DeleteOrder(os.Getenv("AVANZA_ACCOUNT_ID"), res.OrderID)
}
