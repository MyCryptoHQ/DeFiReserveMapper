package process

import (
	"math/big"
	"time"
	"github.com/mycryptohq/DeFiReserveMapper/pkg"
	"github.com/mycryptohq/DeFiReserveMapper/pkg/client"
)

func ProcessAssets(assetItems []root.ImportItem) (map[string]root.ReserveExchangeRate, error) {
	client := client.MakeETHClient()
	outputMap := make(map[string]root.ReserveExchangeRate)
	var compoundItems []root.ImportItem 
	for _, item := range assetItems {
		var returnItems []root.ReserveExchangeRatesObject
		switch item.Type {
			case "uniswap":
				poolTotalReserveSupply, err := FetchPoolTotalSupply(client, item)
				if err == nil || poolTotalReserveSupply == big.NewInt(0) {
					uniswapETHRate, err := BuildUniswapETHReserveRate(client, item, poolTotalReserveSupply)
					if err == nil {
						rateItem :=  root.ReserveExchangeRatesObject{
							AssetId: root.EtherUUID,
							Rate: uniswapETHRate,
						}
						returnItems = append(returnItems, rateItem)
					}
					uniswapERC20Rate, err := BuildUniswapERC20ReserveRate(client, item, poolTotalReserveSupply)
					if err == nil {
						secondRateItem :=  root.ReserveExchangeRatesObject{
							AssetId: item.ReserveTokenUuid,
							Rate: uniswapERC20Rate,
						}
						returnItems = append(returnItems, secondRateItem)
					}
				}
			case "compound":
				compoundItems = append(compoundItems, item)
		}
		if len(returnItems) == 2 {
			resultantObject := root.ReserveExchangeRate{
				Type: item.Type,
				LastUpdated: time.Now().Unix(),
				ReserveRates: returnItems,
			}
			outputMap[item.PoolTokenUuid] = resultantObject
		}
	}

	compoundOutputArr, err := BuildCompoundRates(compoundItems)
	if err == nil {
		for _, compoundRateItem := range compoundOutputArr {
			outputMap[compoundRateItem.PoolTokenUuid] = root.ReserveExchangeRate{
				Type: "compound",
				LastUpdated: time.Now().Unix(),
				ReserveRates: []root.ReserveExchangeRatesObject{{
					AssetId: compoundRateItem.ReserveTokenUuid,
					Rate: compoundRateItem.Rate,
				}},
			}
		}
	}

	return outputMap, nil
}