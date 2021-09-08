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
	api := goavanza.NewApi(logger)

	err := api.Authenticate()
	if err != nil {
		logger.Panic(err)
	}

	api.Listen()
}
