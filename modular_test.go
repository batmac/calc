package calc

import (
	"testing"
)

var tableTestPowm = []struct {
	b, e, m  uint64
	expected uint64
}{
	{1, 1, 2, 1},
	{2, 10, 10, 4},
	{2, 32, 2, 0},
	{2, 64, 2, 0},
	{4, 13, 497, 445},
	{5, 117, 19, 1},
	{2, 4503599627370489, 36028797018963913, 1},
	// {4, 13, math.MaxUint32, 445},
}

func TestPowm(t *testing.T) {
	for _, tt := range tableTestPowm {
		// fmt.Printf("%v^%v (mod %v)\n", tt.b, tt.e, tt.m)
		actual := Powm(tt.b, tt.e, tt.m)
		if actual != tt.expected {
			t.Errorf("Powm(%v, %v, %v) got %v, expected %v", tt.b, tt.e, tt.m, actual, tt.expected)
		}
	}
}

var tableTestMulm = []struct {
	x, b, m  uint64
	expected uint64
}{
	{0, 0, 1, 0},
	{0, 2, 2, 0},
	{2, 0, 2, 0},
	{1, 1, 2, 1},
	{9, 9, 80, 1},
	{9223372036854775807, 9223372036854775807, 100000000000, 84232501249},
}

func TestMulm(t *testing.T) {
	for _, tt := range tableTestMulm {
		actual := Mulm(tt.x, tt.b, tt.m)
		if actual != tt.expected {
			t.Errorf("Mulm(%v,%v,%v) got %v, expected %v", tt.x, tt.b, tt.m, actual, tt.expected)
		}
	}
}
