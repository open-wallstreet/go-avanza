package goavanza

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/monaco-io/request"
)

type TransactionOptions struct {
	From         time.Time
	To           time.Time
	MaxAmount    int
	MinAmount    int
	OrderBookIDs []string
}

func (a *Client) GetTransactions(accountOrTransactionType string, options TransactionOptions) (*Transactions, error) {
	if !a.IsAuthenticated {
		return nil, fmt.Errorf("not authenticated")
	}

	query := make(map[string]string)

	if !options.From.IsZero() {
		query["from"] = options.From.Format(layoutISO)
	}
	if !options.To.IsZero() {
		query["to"] = options.To.Format(layoutISO)
	}
	if len(options.OrderBookIDs) > 0 {
		query["orderbookId"] = strings.Join(options.OrderBookIDs, ",")
	}
	if options.MinAmount > 0 {
		query["minAmount"] = fmt.Sprint(options.MinAmount)
	}
	if options.MaxAmount > 0 {
		query["maxAmount"] = fmt.Sprint(options.MaxAmount)
	}
	a.logger.Infof("%v", query)
	body, _, err := a.request(fmt.Sprintf("/_mobile/account/transactions/%s", accountOrTransactionType), request.GET, nil, RequestOptions{
		query: query,
	})
	if err != nil {
		return nil, err
	}
	var transactions Transactions
	if err := json.Unmarshal([]byte(body), &transactions); err != nil {
		return nil, err
	}
	return &transactions, nil
}
