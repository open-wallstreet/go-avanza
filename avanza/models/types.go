package models

import (
	"encoding/json"
	"strconv"
	"time"
)

type InstrumentType string

const (
	AnyInstrument      InstrumentType = "ANY"
	Stock              InstrumentType = "STOCK"
	Fund               InstrumentType = "FUND"
	Bond               InstrumentType = "BOND"
	Option             InstrumentType = "OPTION"
	FutureForward      InstrumentType = "FUTURE_FORWARD"
	Certificate        InstrumentType = "CERTIFICATE"
	Warrent            InstrumentType = "WARRENT"
	ETF                InstrumentType = "EXCHANGE_TRADED_FUND"
	Index              InstrumentType = "INDEX"
	PremiumBond        InstrumentType = "PREMIUM_BOND"
	SubscriptionOption InstrumentType = "SUBSCRIPTION_OPTION"
	EquityLinkedBond   InstrumentType = "EQUITY_LINKED_BOND"
	Convertible        InstrumentType = "CONVERTIBLE"
)

type OrderType string

const (
	OrderTypeBuy  OrderType = "BUY"
	OrderTypeSell OrderType = "SELL"
)

type TransactionType string

const (
	WITHDRAW    TransactionType = "WITHDRAW"
	DEPOSIT     TransactionType = "DEPOSIT"
	BUY         TransactionType = "BUY"
	SELL        TransactionType = "SELL"
	DIVIDEND    TransactionType = "DIVIDEND"
	DividendTax TransactionType = "DIVIDEND_TAX"
	ForeignTax  TransactionType = "FOREIGN_TAX"
)

// Time represents a long date string of the following format: "2006-01-02T15:04:05.000Z".
type Time time.Time

func (t *Time) UnmarshalJSON(data []byte) error {
	unquoteData, err := strconv.Unquote(string(data))
	if err != nil {
		return err
	}
	parsedTime, err := time.Parse("2006-01-02T15:04:05.000Z", unquoteData)
	if err != nil {
		return err
	}
	*t = Time(parsedTime)
	return nil
}

func (t *Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(*t).Format("2006-01-02T15:04:05.000Z"))
}

// Date represents a short date string of the following format: "2006-01-02".
type Date time.Time

func (d *Date) UnmarshalJSON(data []byte) error {
	unquoteData, err := strconv.Unquote(string(data))
	if err != nil {
		return err
	}
	t, err := time.Parse("2006-01-02", unquoteData)
	if err != nil {
		return err
	}
	*d = Date(t)
	return nil
}

func (d *Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(*d).Format("2006-01-02"))
}

// Millis represents a Unix time in milliseconds since January 1, 1970 UTC.
type Millis time.Time

func (m *Millis) UnmarshalJSON(data []byte) error {
	d, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	*m = Millis(time.UnixMilli(d))
	return nil
}

func (m *Millis) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(*m).UnixMilli())
}

func (m Millis) ToTime() time.Time {
	return time.Time(m)
}

// Nanos represents a Unix time in nanoseconds since January 1, 1970 UTC.
type Nanos time.Time

func (n *Nanos) UnmarshalJSON(data []byte) error {
	d, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}

	// Go Time package does not include a method to convert UnixNano to a time.
	timeNano := time.Unix(d/1_000_000_000, d%1_000_000_000)
	*n = Nanos(timeNano)
	return nil
}

func (n *Nanos) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(*n).UnixNano())
}

func (n Nanos) ToTime() time.Time {
	return time.Time(n)
}
