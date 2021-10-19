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

	err := api.Authenticate()
	if err != nil {
		logger.Panic(err)
	}

	options := goavanza.NewAvanzaWebsocketOptions()
	options.OnError = func(e error) {
		logger.Errorf("websocket error: %v", err)
	}
	options.OnConnected = func() {
		logger.Info("websocket connected")
	}
	options.OnDisconnect = func(e error) {
		logger.Infof("websocket disconnected", err)
	}
	options.OnQuote = func(quote goavanza.QuoteMessage) {
		logger.Info(quote)
	}
	options.OnOrderDepthsMessage = func(od goavanza.OrderDepthsMessage) {
		logger.Info(od)
	}

	websocketApi := goavanza.NewWebsocket(api, logger, options)

	websocketApi.Subscribe([]string{"/orderdepths/549768"})
	websocketApi.Listen()
}
