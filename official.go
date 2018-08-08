package prime

import (
	"math/big"
)

func Official(n int64) bool {
	return big.NewInt(n).ProbablyPrime(0)
}