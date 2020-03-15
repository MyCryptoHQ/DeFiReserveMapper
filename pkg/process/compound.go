package process

import (
	"strings"
	"math/big"
	"github.com/mycryptohq/DeFiReserveMapper/pkg"
	"github.com/mycryptohq/DeFiReserveMapper/pkg/compoundapi"
)

type CompoundRate struct {
	PoolTokenUuid			string			`json:"poolTokenUuid"`
	ReserveTokenUuid		string			`json:"reserveTokenUuid"`
	Rate					*big.Float		`json:"rate"`
}

func BuildCompoundRates(compoundItems []root.ImportItem) ([]CompoundRate, error) {
	var compoundReturnArray []CompoundRate
	compoundClient := compoundapi.MakeCompoundApiClient()
	
	cTokens, err := compoundapi.FetchCompoundCTokens(compoundClient)
	if err != nil {
		return []CompoundRate{}, nil
	}
	for _, compoundItem := range compoundItems {
		var relevantCToken compoundapi.CToken
		for _, ctoken := range cTokens {
			if strings.ToLower(ctoken.TokenAddress) == strings.ToLower(compoundItem.PoolTokenAddress) {
				relevantCToken = ctoken
				break;
			}
		}
		exchangeRate, success := new(big.Float).SetString(relevantCToken.ExchangeRate.Value)
		if success {
			compoundReturnArray = append(compoundReturnArray, CompoundRate{
				PoolTokenUuid: compoundItem.PoolTokenUuid,
				ReserveTokenUuid: compoundItem.ReserveTokenUuid,
				Rate: exchangeRate,
			})
		}
	} 
	return compoundReturnArray, nil
}