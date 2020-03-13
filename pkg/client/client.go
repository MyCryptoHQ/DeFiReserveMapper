package client

import (
	"fmt"
	"log"
	"github.com/mycryptohq/DeFiReserveMapper/pkg"
	//"github.com/mycryptohq/DeFiReserveMapper/pkg/process"
	//"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func MakeETHClient () (*ethclient.Client) {
	configEndpoint := root.NodeEndpoint
	fmt.Println("Starting to make eth client")
	client, err := ethclient.Dial(configEndpoint)
	if err != nil {
		log.Fatalf("Could not collect to eth client")
	}
	return client
}


func getBalance(client *ethclient.Client, address string) (error) {
	return nil
}