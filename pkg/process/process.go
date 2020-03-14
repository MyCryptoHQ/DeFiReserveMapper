package process

import (
	"time"
	"github.com/mycryptohq/DeFiReserveMapper/pkg"
	"github.com/mycryptohq/DeFiReserveMapper/pkg/client"
)

func ProcessAssets(assetItems []root.ImportItem) (map[string]root.ReserveExchangeRate, error) {
	client := client.MakeETHClient()
	outputMap := make(map[string]root.ReserveExchangeRate)
	for _, item := range assetItems {
		var returnItems []root.ReserveExchangeRatesObject
		switch item.Type {
		case "uniswap":
			uniswapETHRate, err := BuildUniswapETHReserveRate(client, item)
			if err == nil {
				rateItem :=  root.ReserveExchangeRatesObject{
					AssetId: root.EtherUUID,
					Rate: uniswapETHRate,
				}
				returnItems = append(returnItems, rateItem)
			}

			uniswapERC20Rate, err := BuildUniswapERC20ReserveRate(client, item)
			if err == nil {
				secondRateItem :=  root.ReserveExchangeRatesObject{
					AssetId: item.ReserveTokenUuid,
					Rate: uniswapERC20Rate,
				}
				returnItems = append(returnItems, secondRateItem)
			}	
		}
		resultantObject := root.ReserveExchangeRate{
			Type: item.Type,
			LastUpdated: time.Now().Unix(),
			ReserveRates: returnItems,
		}
		outputMap[item.PoolTokenUuid] = resultantObject
	}
	return outputMap, nil
}