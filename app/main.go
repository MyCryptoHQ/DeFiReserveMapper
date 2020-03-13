package main

import (
	"log"
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
	"github.com/mycryptohq/DeFiReserveMapper/pkg"
	"github.com/mycryptohq/DeFiReserveMapper/pkg/process"
)

func main() {
	fmt.Println("Starting pull")
	jsonFile, err := os.Open("../test-assets.json")
	if err != nil {
		log.Fatalf("Couldn't open file")
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var itemsToProcess []root.ImportItem

	json.Unmarshal(byteValue, &itemsToProcess)
	if len(itemsToProcess) == 0 {
		return
	}
	results, err := process.ProcessAssets(itemsToProcess)
	if err != nil {
		return
	}
	fmt.Println(results)
}