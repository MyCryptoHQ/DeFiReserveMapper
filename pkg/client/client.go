package client

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mycryptohq/DeFiReserveMapper/pkg"
	"log"
)

func MakeETHClient() *ethclient.Client {
	configEndpoint := root.NodeEndpoint
	client, err := ethclient.Dial(configEndpoint)
	if err != nil {
		log.Fatalf("Could not collect to eth client")
	}
	return client
}
