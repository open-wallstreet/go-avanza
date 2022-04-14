package main

import (
	"github.com/joho/godotenv"
	goavanza "github.com/open-wallstreet/go-avanza"
	"go.uber.org/zap"
)

func main() {
	godotenv.Load()
	log, _ := zap.NewDevelopment()
	logger := log.Sugar()
	api := goavanza.NewClient(logger)
	api.Authenticate()
	search, err := api.Search("novotek", goavanza.STOCK)
	if err != nil {
		panic(err)
	}
	logger.Info(search)

	o, err := api.GetOrderbook(search.Hits[0].TopHits[0].ID, goavanza.STOCK)
	if err != nil {
		panic(err)
	}
	logger.Info(o)

}
