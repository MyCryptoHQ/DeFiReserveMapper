package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/mycryptohq/DeFiReserveMapper/pkg"
	"github.com/mycryptohq/DeFiReserveMapper/pkg/process"
	"github.com/mycryptohq/DeFiReserveMapper/pkg/s3"
)

var (
	bucket = os.Getenv("bucket")
	region = os.Getenv("region")
)

func handleRequest() {
	// Fetch input file
	jsonFile, err := os.Open("./assets.json")
	if err != nil {
		log.Fatal("Couldn't open assets config file", err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var allItemsToProcess []root.ImportItem

	json.Unmarshal(byteValue, &allItemsToProcess)
	if len(allItemsToProcess) == 0 {
		log.Fatalf("No items to process found")
	}
	// Fetch output file
	outputItems := make(map[string]root.ReserveExchangeRate)

	outputFile := "/tmp/outputFile.json"

	// Download output file from s3
	if err := s3.Download(bucket, region, outputFile); err != nil {
		log.Println("\nCouldn't open output file. Assume this is first run.")
	} else {
		outputFile, err := os.Open(outputFile)
		if err != nil {
			log.Printf("Couldn't open %s", outputFile.Name())
		}
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

	// Merge the results into old file, overwriting updated values
	mergedResults := mergeChanges(results, outputItems)

	fmt.Printf("Updated %d pool tokens.\n", len(filteredItemsToProcess))
	file, _ := json.MarshalIndent(mergedResults, "", "    ")

	// Upload to s3
	if err := s3.Upload(bucket, region, outputFile, bytes.NewReader(file)); err != nil {
		log.Println("Error uploading to s3", err)
	}
}
func main() {
	lambda.Start(handleRequest)
}

func filterItemsToProcess(allInputItems []root.ImportItem, outputFileObject map[string]root.ReserveExchangeRate) []root.ImportItem {
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
		if !foundItem || (currentTime-outputItem.LastUpdated >= updateInterval) {
			outputItemArray = append(outputItemArray, item)
		}
	}
	return outputItemArray
}

func mergeChanges(fetchedRateObject map[string]root.ReserveExchangeRate, outputFileObject map[string]root.ReserveExchangeRate) map[string]root.ReserveExchangeRate {
	keysArray := make([]string, len(fetchedRateObject))
	i := 0
	// Makes an array of keys from the fetched rate object
	for k := range fetchedRateObject {
		keysArray[i] = k
		i++
	}

	// for each key (pool token uuid), overwrite the existing entry in the previously-imported asset file
	for _, key := range keysArray {
		outputFileObject[key] = fetchedRateObject[key]
	}
	return outputFileObject
}
