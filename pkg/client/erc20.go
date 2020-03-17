package client

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func BalanceOf(client *ethclient.Client, tokenObject TokenBalance) (*big.Int, error) {

	tokenCaller, err := NewTokenCaller(tokenObject.Contract, client)
	if err != nil {
		return big.NewInt(0), err
	}

	balance, err := tokenCaller.BalanceOf(nil, tokenObject.Wallet)
	if err != nil {
		return big.NewInt(0), err
	}
	return balance, nil
}

func TotalSupply(client *ethclient.Client, tokenObject TokenBalance) (*big.Int, error) {

	tokenCaller, err := NewTokenCaller(tokenObject.Contract, client)
	if err != nil {
		return big.NewInt(0), err
	}

	totalSupply, err := tokenCaller.TotalSupply(nil)
	if err != nil {
		return big.NewInt(0), err
	}
	return totalSupply, nil
}
