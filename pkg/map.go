package root

import (
	"math/big"
)

type ReserveExchangeRate struct {
	Type					string							`json:"type"`
	ReserveRates			[]ReserveExchangeRatesObject	`json:"reserveRates"`
}

type ReserveExchangeRatesObject struct {
	AssetId			string							`json:"assetId"`
	Rate			*big.Float						`json:"rate"`
}

type ImportItem struct {
	Type					string		`json:"type"` 	// uniswap || compound
	Name					string		`json:"name"`
	ReserveTokenAddress		string		`json:"reserveTokenAddress"` // reserve token contract address
	ReserveTokenDecimals	int			`json:"reserveTokenDecimals"`
	ReserveTokenUuid		string		`json:"reserveTokenUuid"`
	PoolTokenUuid			string		`json:"poolTokenUuid"`
	PoolTokenAddress		string		`json:"poolTokenAddress"` // pool token contract address
	PoolTokenDecimals		int			`json:"poolTokenDecimals"`
}