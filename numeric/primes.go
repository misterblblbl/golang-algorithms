package numeric

// IsPrime is the implementation of algorithm with complexity O(n^(0.5))
func IsPrime(num int) bool {
	if num == 1 {
		return false
	}

	i := 2

	for (i * i) < num {
		if num%i == 0 {
			return false
		}
		i++
	}

	return true
}
