package main

import (
	"time"
	"log"
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
	"github.com/mycryptohq/DeFiReserveMapper/pkg"
	"github.com/mycryptohq/DeFiReserveMapper/pkg/process"
)

func main() {

	// Fetch input file
	jsonFile, err := os.Open("../test-assets.json")
	if err != nil {
		log.Fatalf("Couldn't open assets config file")
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var allItemsToProcess []root.ImportItem

	json.Unmarshal(byteValue, &allItemsToProcess)
	if len(allItemsToProcess) == 0 {
		return
	}

	// Fetch output file
	outputItems := make(map[string]root.ReserveExchangeRate)
	outputFile, err := os.Open("../outputFile.json")
	if err != nil {
		fmt.Println("Couldn't open output file. Assume this is first run.")
	} else {
		defer outputFile.Close()
		byteOutputFileValue, _ := ioutil.ReadAll(outputFile)
		json.Unmarshal(byteOutputFileValue, &outputItems)
	}

	// Filter to only include assets needed
	filteredItemsToProcess := filterItemsToProcess(allItemsToProcess, outputItems)
	results, err := process.ProcessAssets(filteredItemsToProcess)
	if err != nil {
		return
	}

	mergedResults := mergeChanges(results, outputItems)
	fmt.Printf("Updated %d pool tokens.\n", len(filteredItemsToProcess))
	file, _ := json.MarshalIndent(mergedResults, "", "    ")
 
	_ = ioutil.WriteFile("../outputFile.json", file, 0644)
}

func filterItemsToProcess(allInputItems []root.ImportItem, outputFileObject map[string]root.ReserveExchangeRate) ([]root.ImportItem) {
	var outputItemArray []root.ImportItem
	for _, item := range allInputItems {
		var updateInterval int64
		currentTime := time.Now().Unix()
		if interval := item.UpdateInterval; interval != 0 {
			updateInterval = int64(interval)
		} else {
			updateInterval = int64(root.DefaultRefreshInterval)
		}
		outputItem, foundItem := outputFileObject[item.PoolTokenUuid]
		if !foundItem || (currentTime - outputItem.LastUpdated >= updateInterval) {
			outputItemArray = append(outputItemArray, item)
		}
	}
	return outputItemArray
}

func mergeChanges(fetchedRateObject map[string]root.ReserveExchangeRate, outputFileObject map[string]root.ReserveExchangeRate) (map[string]root.ReserveExchangeRate) {
	keysArray := make([]string, len(fetchedRateObject))

	i := 0
	for k := range fetchedRateObject {
		keysArray[i] = k
		i++
	}

	for _, key := range keysArray {
		outputFileObject[key] = fetchedRateObject[key]
	}
	return outputFileObject
}