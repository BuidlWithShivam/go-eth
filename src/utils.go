package src

import (
	"math"
	"math/big"
)

func BigIntToEthBigFloat(balance *big.Int, decimals int) *big.Float {
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	return new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(decimals)))
}
