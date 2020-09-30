package simple

import (
	"math"
)

func reverse(x int) int {
	max := int64(math.Abs(math.Pow(float64(2), 31))) - 1
	min := -1 * (max + 1)
	if x == 0 {
		return 0
	}

	absX := int64(math.Abs(float64(x)))
	absXRoot := absX

	var result int64
	for index := 0; absX >= 1; index++ {
		if result*10+absX%10 > max || result*10+absX%10*-1 < min {
			return 0
		} else {
			result = result*10 + absX%10
			absX = absX / 10
		}
	}
	result = result * int64(x) / absXRoot

	return int(result)
}
