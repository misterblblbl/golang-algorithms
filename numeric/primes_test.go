package numeric

import "testing"

type testCase struct {
	number  int
	isPrime bool
}

var cases = []testCase{
	{3, true},
	{5, true},
	{12, false},
	{122, false},
	{113, true},
	{123, false},
	{1000000, false},
	{802597, true},
}

func TestIsPrime(t *testing.T) {
	for _, tc := range cases {
		got := IsPrime(tc.number)

		if got != tc.isPrime {
			t.Errorf("Number %d: expected %v, got %v", tc.number, tc.isPrime, got)
		}
	}
}
