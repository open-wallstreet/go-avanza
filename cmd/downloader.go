package cmd

import (
	"context"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/jszwec/csvutil"
	client2 "github.com/open-wallstreet/go-avanza/avanza/client"
	"github.com/open-wallstreet/go-avanza/avanza/market"
	"github.com/open-wallstreet/go-avanza/avanza/models"
	"github.com/schollz/progressbar/v3"

	"github.com/spf13/cobra"
)

// downloaderCmd represents the downloader command
var downloaderCmd = &cobra.Command{
	Use:   "downloader",
	Short: "Extract and download data into files",
	Long: `Extract and download different types of data from Avanza that can be saved into files
for easy import into databases or other programs.`,
}

type stocksListData struct {
	AvanzaID              string `csv:"avanza_id"`
	ISIN                  string `csv:"isin"`
	MarketCapital         int    `csv:"market_capital"`
	MarketCapitalCurrency string `csv:"market_capital_currency"`
	TotalNumberOfShares   int    `csv:"total_number_of_shares"`
	Country               string `csv:"country"`
	Name                  string `csv:"name"`
	MarketList            string `csv:"market_list"`
	Ticker                string `csv:"ticker"`
	Sector                string `csv:"sector"`
	Type                  string `csv:"type"`
}

var stocksListCmd = &cobra.Command{
	Use:   "stocks-list",
	Short: "Extract and download list of available stocks",
	Long:  `Downloads all available stocks available on Avanza.se.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("stocks-list called")
		fmt.Println(OutputPath)
		ids := extractAvailableIds()
		sizeOfIds := len(ids)
		bar := progressbar.Default(int64(sizeOfIds))

		client := market.MarketClient{
			Client: client2.New(),
		}
		data := make([]*stocksListData, 0, sizeOfIds)
		for _, id := range ids {
			bar.Add(1)
			instrument, err := client.GetInstrument(context.Background(), &models.GetInstrumentParams{
				Instrument: models.Stock,
				ID:         id,
			})
			if err != nil {
				fmt.Println(err)
				return
			}
			data = append(data, &stocksListData{
				AvanzaID:              instrument.OrderbookID,
				ISIN:                  instrument.Isin,
				MarketCapital:         int(instrument.KeyIndicators.MarketCapital.Value),
				MarketCapitalCurrency: instrument.Listing.Currency,
				Country:               instrument.Listing.CountryCode,
				Name:                  instrument.Name,
				MarketList:            instrument.Listing.MarketListName,
				Ticker:                instrument.Listing.TickerSymbol,
				Type:                  instrument.Type,
			})
		}

		b, err := csvutil.Marshal(data)
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		permissions := 0644 // or whatever you need
		err = ioutil.WriteFile(OutputPath, b, fs.FileMode(permissions))
		if err != nil {
			fmt.Println("error:", err)
			return
		}
	},
}

var OutputPath = ""

func init() {
	rootCmd.AddCommand(downloaderCmd)
	downloaderCmd.AddCommand(stocksListCmd)
	stocksListCmd.Flags().StringVarP(&OutputPath, "output", "o", "", "output file")
	stocksListCmd.MarkFlagRequired("output")
}

const ExtractUrl = "https://www.avanza.se/frontend/template.html/marketing/advanced-filter/advanced-filter-template?%d&parameters.startIndex=0&parameters.maxResults=90000"

func extractAvailableIds() []string {
	url := fmt.Sprintf(ExtractUrl, time.Now().UnixMilli())
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var re = regexp.MustCompile(`(?m)/aktier/om-aktien.html/(\d+)/.*`)
	var ids []string
	doc.Find(".row .orderbookName").Each(func(i int, s *goquery.Selection) {
		node := s.Find("a")
		url, ok := node.Attr("href")
		if !ok {
			fmt.Println("failed to find url in node")
			return
		}
		submatch := re.FindStringSubmatch(url)
		if len(submatch) < 1 {
			fmt.Println("failed to find id")
			return
		}
		id := submatch[1]
		ids = append(ids, strings.TrimSpace(id))
	})
	return ids
}
