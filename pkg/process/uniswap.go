package process

import (
	"math"
	"math/big"
	"github.com/mycryptohq/DeFiReserveMapper/pkg"
	"github.com/mycryptohq/DeFiReserveMapper/pkg/client"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
)

func BuildUniswapETHReserveRate(ETHClient *ethclient.Client, importedItem root.ImportItem, poolReserveTotalSupply *big.Int) (*big.Float, error) {
	poolReserveETH, err := client.GetBalance(*ETHClient, importedItem.PoolTokenAddress)
	if err != nil {
		return nil, err
	}

	floatPoolReserveETH := new(big.Float).SetInt(poolReserveETH)
	floatPoolTotalSupply := new(big.Float).SetInt(poolReserveTotalSupply)
	precision, _ := new(big.Float).Quo(big.NewFloat(1), big.NewFloat(math.Pow10(int(18)))).Float64()
	outputRate, _ := new(big.Float).Quo(floatPoolReserveETH, floatPoolTotalSupply).Float64()
	truncatedRate := Truncate(outputRate, precision)
	return big.NewFloat(truncatedRate), nil
}

func BuildUniswapERC20ReserveRate(ETHClient *ethclient.Client, importedItem root.ImportItem, poolReserveTotalSupply *big.Int) (*big.Float, error) {

	tokenObject := client.TokenBalance{
		Contract: common.HexToAddress(importedItem.ReserveTokenAddress),
		Wallet: common.HexToAddress(importedItem.PoolTokenAddress),
		Name: importedItem.Name,
		Decimals: int64(importedItem.ReserveTokenDecimals),
	}
	poolReserveERC20, err := client.BalanceOf(ETHClient, tokenObject)
	if err != nil {
		return nil, err
	}
	
	floatPoolReserveERC20 := new(big.Float).SetInt(poolReserveERC20)
	floatPoolTotalSupply := new(big.Float).SetInt(poolReserveTotalSupply)

	multiplier := big.NewFloat(math.Pow10(int(18-importedItem.ReserveTokenDecimals)))
	precision, _ := new(big.Float).Quo(big.NewFloat(1), big.NewFloat(math.Pow10(int(18)))).Float64()
	
	outputRateBase := new(big.Float).Quo(floatPoolReserveERC20, floatPoolTotalSupply)
	outputRate, _ := new(big.Float).Mul(multiplier, outputRateBase).Float64()
	truncatedRate := Truncate(outputRate, precision)
	return big.NewFloat(truncatedRate), nil
}

func FetchPoolTotalSupply(ETHClient *ethclient.Client, importedItem root.ImportItem) (*big.Int, error) {
	poolTokenObject := client.TokenBalance{
		Contract: common.HexToAddress(importedItem.PoolTokenAddress),
		Name: importedItem.Name,
		Decimals: int64(importedItem.PoolTokenDecimals),
	}

	poolReserveTotalSupply, err := client.TotalSupply(ETHClient, poolTokenObject)
	if err != nil {
		return nil, err
	}
	return poolReserveTotalSupply, nil
}