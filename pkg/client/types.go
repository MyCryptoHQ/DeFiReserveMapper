package client

import (
	"context"
	"math/big"
	"github.com/ethereum/go-ethereum/common"
)

type TokenBalance struct {
	Contract common.Address
	Wallet   common.Address
	Name     string
	Symbol   string
	Balance  *big.Int
	ETH      *big.Int
	Decimals int64
	Block    int64
	ctx      context.Context
}