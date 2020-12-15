package numeric

import (
	"math"
)

func Factorize(num int) []int {
	factors := make([]int, 0, int(math.Sqrt(float64(num))))

	f := 2
	for num%f == 0 {
		factors = append(factors, f)
		num = num / f
	}

	f = 3
	for f*f <= num {
		if num%f == 0 {
			factors = append(factors, f)
			num = num / f
		} else {
			f = f + 2
		}
	}

	if num != 1 {
		factors = append(factors, num)
	}

	return factors
}
