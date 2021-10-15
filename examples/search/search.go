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

	search, err := api.Search("novotek", goavanza.STOCK)
	if err != nil {
		panic(err)
	}

	logger.Info(search)
}
