package goavanza

import (
	"encoding/json"
	"fmt"
	"github.com/monaco-io/request"
	"strings"
)

type Instrument string

const (
	ANY                 Instrument = "ANY"
	STOCK               Instrument = "STOCK"
	FUND                Instrument = "FUND"
	BOND                Instrument = "BOND"
	OPTION              Instrument = "OPTION"
	FUTURE_FORWARD      Instrument = "FUTURE_FORWARD"
	CERTIFICATE         Instrument = "CERTIFICATE"
	WARRENT             Instrument = "WARRENT"
	ETF                 Instrument = "EXCHANGE_TRADED_FUND"
	INDEX               Instrument = "INDEX"
	PREMIUM_BOND        Instrument = "PREMIUM_BOND"
	SUBSCRIPTION_OPTION Instrument = "SUBSCRIPTION_OPTION"
	EQUITY_LINKED_BOND  Instrument = "EQUITY_LINKED_BOND"
	CONVERTIBLE         Instrument = "CONVERTIBLE"
)

func (e Instrument) String() string {
	instruments := [...]string{
		"STOCK",
		"FUND",
		"BOND",
		"OPTION",
		"FUTURE_FORWARD",
		"CERTIFICATE",
		"WARRENT",
		"EXCHANGE_TRADED_FUND",
		"INDEX",
		"PREMIUM_BOND",
		"SUBSCRIPTION_OPTION",
		"EQUITY_LINKED_BOND",
		"CONVERTIBLE",
	}

	x := string(e)
	for _, v := range instruments {
		if v == x {
			return x
		}
	}

	return ""
}

func (a *Client) GetInstrument(id string, instrument Instrument) (*InstrumentResponse, error) {
	body, _, err := a.request(fmt.Sprintf("/_mobile/market/%s/%s", strings.ToLower(string(instrument)), id), request.GET, nil, RequestOptions{})
	if err != nil {
		return nil, err
	}
	var instrumentResponse InstrumentResponse
	if err := json.Unmarshal([]byte(body), &instrumentResponse); err != nil {
		return nil, err
	}

	return &instrumentResponse, nil
}
