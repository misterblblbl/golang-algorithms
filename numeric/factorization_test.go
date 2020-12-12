package numeric

import "testing"

type testFactors struct {
	number  int
	factors []int
}

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

var casesFactors = []testFactors{
	{3, []int{3}},
	{5, []int{5}},
	{12, []int{2, 2, 3}},
	{122, []int{2, 61}},
	{113, []int{113}},
	{123, []int{3, 41}},
	{1000000, []int{2, 2, 2, 2, 2, 2, 5, 5, 5, 5, 5, 5}},
	{802597, []int{802597}},
	{917521579, []int{13, 70578583}},
}

func TestFactoriztion(t *testing.T) {
	for _, cf := range casesFactors {
		got := Factorize(cf.number)
		eq := slicesEqual(cf.factors, got)

		if !eq {
			t.Errorf("Number %d: expected %v, got %v", cf.number, cf.factors, got)
		}
	}
}

var result []int

func BenchmarkFactorization(b *testing.B) {
	var r []int
	difficultCase := 917521579

	for i := 0; i < b.N; i++ {
		// always record the result of Fib to prevent
		// the compiler eliminating the function call.
		r = Factorize(difficultCase)
	}

	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}
