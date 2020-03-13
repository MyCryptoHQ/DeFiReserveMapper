package process

import (
	"fmt"
	"github.com/mycryptohq/DeFiReserveMapper/pkg"
	"github.com/mycryptohq/DeFiReserveMapper/pkg/client"
)

func ProcessAssets(assetItems []root.ImportItem) ([]root.ReserveExchangeRatesObject, error) {
	client := client.MakeETHClient()
	var returnItems []root.ReserveExchangeRatesObject
	for _, item := range assetItems {
		switch item.Type {
		case "uniswap":
			uniswapETHRate, err := BuildUniswapETHReserveRate(client, item)
			if err != nil {
				fmt.Println(err)
				return returnItems, err
			}
			
			rateItem :=  root.ReserveExchangeRatesObject{
				AssetId: root.EtherUUID,
				Rate: uniswapETHRate,
			}
			returnItems = append(returnItems, rateItem)

			uniswapERC20Rate, err := BuildUniswapERC20ReserveRate(client, item)
			if err != nil {
				fmt.Println(err)
				return returnItems, err
			}
			
			secondRateItem :=  root.ReserveExchangeRatesObject{
				AssetId: item.ReserveTokenUuid,
				Rate: uniswapERC20Rate,
			}
			returnItems = append(returnItems, secondRateItem)
		}
	}
	return returnItems, nil
}