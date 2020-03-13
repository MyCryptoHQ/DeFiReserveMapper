package root

import (
	"math/big"
)

type ReserveExchangeRate struct {
	Rates			[]ReserveExchangeRatesObject	`json:"rates"`
}

type ReserveExchangeRatesObject struct {
	Type			string							`json:"type"`
	Rate			*big.Float						`json:"rate"`
}

type ImportItem struct {
	Uuid				string		`json:"uuid"`
	Type				string		`json:"type"` 	// uniswap || compound
	Name				string		`json:"name"`
	ContractAddress		string		`json:"contractAddress"`
}