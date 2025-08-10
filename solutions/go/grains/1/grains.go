package grains

import (
	"errors"
	"math"
)

const cheeseboardTotalGrains = 18446744073709551615

func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, errors.New("number must be larger than 0 and smaller than or equal to 64")
	}
	return uint64(math.Pow(2, float64(number-1))), nil
}

func Total() uint64 {
	return cheeseboardTotalGrains
}
