package process

import (
	"fmt"
	"math"
	"math/big"
	"github.com/mycryptohq/DeFiReserveMapper/pkg"
	"github.com/mycryptohq/DeFiReserveMapper/pkg/client"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
)

func BuildUniswapETHReserveRate(ETHClient *ethclient.Client, importedItem root.ImportItem) (*big.Float, error) {
	poolReserveETH, err := client.GetBalance(*ETHClient, importedItem.PoolTokenAddress)
	if err != nil {
		return nil, err
	}

	poolTokenObject := client.TokenBalance{
		Contract: common.HexToAddress(importedItem.PoolTokenAddress),
		Name: importedItem.Name,
		Decimals: int64(importedItem.PoolTokenDecimals),
	}

	fmt.Println(importedItem)

	poolReserveTotalSupply, err := client.TotalSupply(ETHClient, poolTokenObject)
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

func BuildUniswapERC20ReserveRate(ETHClient *ethclient.Client, importedItem root.ImportItem) (*big.Float, error) {

	tokenObject := client.TokenBalance{
		Contract: common.HexToAddress(importedItem.ReserveTokenAddress),
		Wallet: common.HexToAddress(importedItem.PoolTokenAddress),
		Name: importedItem.Name,
		Decimals: int64(importedItem.ReserveTokenDecimals),
	}
	poolTokenObject := client.TokenBalance{
		Contract: common.HexToAddress(importedItem.PoolTokenAddress),
		Name: importedItem.Name,
		Decimals: int64(importedItem.PoolTokenDecimals),
	}

	poolReserveTotalSupply, err := client.TotalSupply(ETHClient, poolTokenObject)
	if err != nil {
		return nil, err
	}
	poolReserveERC20, err := client.BalanceOf(ETHClient, tokenObject)
	if err != nil {
		return nil, err
	}
	floatPoolReserveERC20 := new(big.Float).SetInt(poolReserveERC20)
	floatPoolTotalSupply := new(big.Float).SetInt(poolReserveTotalSupply)

	precision, _ := new(big.Float).Quo(big.NewFloat(1), big.NewFloat(math.Pow10(int(18)))).Float64()
	
	outputRate, _ := new(big.Float).Quo(floatPoolReserveERC20, floatPoolTotalSupply).Float64()
	truncatedRate := Truncate(outputRate, precision)
	return big.NewFloat(truncatedRate), nil
}