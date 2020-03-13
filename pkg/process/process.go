package process

import (
	"math/big"
)

func convertToBase(num *big.Int) (error, *big.Int) {
	if num == big.NewInt(0) {
		return nil, num
	}
	return nil, big.NewInt(0)
}