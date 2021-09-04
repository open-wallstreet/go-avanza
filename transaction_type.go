package goavanza

type TransactionType string

const (
	WITHDRAW     TransactionType = "WITHDRAW"
	DEPOSIT      TransactionType = "DEPOSIT"
	BUY          TransactionType = "BUY"
	SELL         TransactionType = "SELL"
	DIVIDEND     TransactionType = "DIVIDEND"
	DIVIDEND_TAX TransactionType = "DIVIDEND_TAX"
	FOREIGN_TAX  TransactionType = "FOREIGN_TAX"
)

func (e TransactionType) String() string {
	instruments := [...]string{
		"WITHDRAW",
		"DEPOSIT",
		"BUY",
		"SELL",
	}

	x := string(e)
	for _, v := range instruments {
		if v == x {
			return x
		}
	}

	return ""
}
