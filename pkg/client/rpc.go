package client

import (
	"context"
	"math/big"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func getBalance(client ethclient.Client, address string) (*big.Int, error) {
	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return big.NewInt(0), err
	}
	return balance, nil
}