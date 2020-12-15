package numeric

import "testing"

type testCaseGcd struct {
	a      int
	b      int
	result int
}

var casesGcd = []testCaseGcd{
	{60, 24, 12},
	{24, 60, 12},
	{60, 60, 60},
	{7, 13, 1},
	{7, 14, 7},
	{3283, 539, 49},
	{0, 539, 539},
	{539, 0, 539},
	{0, 0, 0},
}

func TestGcd(t *testing.T) {
	for _, tc := range casesGcd {
		got := Gcd(tc.a, tc.b)

		if got != tc.result {
			t.Errorf("Number Gcd(%d, %d): expected %d, got %d", tc.a, tc.b, tc.result, got)
		}
	}
}
