package sumdecimal

import (
	"math"
	"math/big"
)

var (
	one     = big.NewInt(1)
	ten     = big.NewInt(10)
	twenty  = big.NewInt(20)
	hundred = big.NewInt(100)
)

func SumDecimal(c int) int {
	if c <= 0 {
		return 0
	}

	root := int(math.Sqrt(float64(c)))
	c = c - (root * root)

	bigC := big.NewInt(int64(c))
	bigX := big.NewInt(0)
	bigS := big.NewInt(0)
	bigM := big.NewInt(0)
	bigP := big.NewInt(int64(root * 20))

	var sum int64

	for i := 0; i < 1000; i++ {
		bigC.Mul(bigC, hundred)
		bigX.DivMod(bigC, bigP, bigM)
		bigS.Mul(bigX, bigX)

		bigC.Set(bigM)
		bigC.Sub(bigC, bigS)

		if bigS.Cmp(bigM) > 0 {
			bigX.Sub(bigX, one)
			bigC.Add(bigC, bigP)
			bigC.Add(bigC, bigX)
			bigC.Add(bigC, bigX)
			bigC.Add(bigC, one)
		}

		sum += bigX.Int64()

		bigX.Mul(bigX, twenty)
		bigP.Mul(bigP, ten)
		bigP.Add(bigP, bigX)
	}

	return int(sum)
}
