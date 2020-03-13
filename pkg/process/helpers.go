package process

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

func Truncate(f float64, unit float64) float64 {
    bf := big.NewFloat(0).SetPrec(1000).SetFloat64(f)
    bu := big.NewFloat(0).SetPrec(1000).SetFloat64(unit)

    bf.Quo(bf, bu)

    // Truncate:
    i := big.NewInt(0)
    bf.Int(i)
    bf.SetInt(i)

    f, _ = bf.Mul(bf, bu).Float64()
    return f
}