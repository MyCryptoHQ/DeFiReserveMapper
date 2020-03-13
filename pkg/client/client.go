package client

import (
	"log"
	"github.com/mycryptohq/DeFiReserveMapper/pkg"
	"github.com/ethereum/go-ethereum/ethclient"
)

func MakeETHClient () (*ethclient.Client) {
	configEndpoint := root.NodeEndpoint
	client, err := ethclient.Dial(configEndpoint)
	if err != nil {
		log.Fatalf("Could not collect to eth client")
	}
	return client
}
