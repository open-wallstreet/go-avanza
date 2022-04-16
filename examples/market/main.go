package main

import (
	"context"
	avanza "github.com/open-wallstreet/go-avanza"
	"github.com/open-wallstreet/go-avanza/avanza/models"
	"log"
)

func main() {
	client := avanza.New(avanza.WithDebug(true))
	defer client.Close()
	search, err := client.Market.Search(context.Background(), &models.SearchParams{
		Query:      "novotek",
	})
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println(search)
	instrument, err := client.Market.GetInstrument(context.Background(), &models.GetInstrumentParams{
		Instrument: search.Hits[0].InstrumentType,
		ID:         search.Hits[0].TopHits[0].ID,
	})
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println(instrument)
}
