package helpers

import (
	"math"
	"math/big"
)

func ConvertFromWei(amount big.Int) float64 {
	inputAmountFloat := new(big.Float).SetInt(&amount)
	output := new(big.Float).Quo(inputAmountFloat, big.NewFloat(math.Pow10(int(18))))
	outputFloat, _ := output.Float64()
	return outputFloat
}

func ConvertFromBase(amount big.Int, decimal int) float64 {
	inputAmountFloat := new(big.Float).SetInt(&amount)
	output := new(big.Float).Quo(inputAmountFloat, big.NewFloat(math.Pow10(int(decimal))))
	outputFloat, _ := output.Float64()
	return outputFloat
}
