package calc

import (
	"testing"
)

var tableTestIsPrime = []struct {
	in       uint64
	expected bool
}{
	{0, false},
	{1, false},
	{2, true},
	{3, true},
	{4, false},
	{5, true},
	{6, false},
	{7, true},
	{8, false},
	{9, false},
	{10, false},
	{2047, false},
	{341550071728321, false},
	{1<<8 - 5, true},
	{1<<18 - 5, true},
	{1<<26 - 5, true},
	{1<<32 - 5, true},
	{1<<36 - 5, true},
	{1<<36 - 17, true},
	{1<<36 - 23, true},
	{1<<40 - 87, true},
	{1<<48 - 59, true},
	{1<<55 - 55, true},
	{1<<61 - 1, true},
	{1<<63 - 25, true},
	// {1<<64 - 363, true},
	{5537838510751, false},
	{18442979630128095121, false},
	//from "Smallest odd number for which Miller-Rabin primality test on bases <= n-th prime does not reveal compositeness."
	{2047, false},
	{1373653, false},
	{25326001, false},
	{3215031751, false},
	{2152302898747, false},
	{3474749660383, false},
	{341550071728321, false},
	{3825123056546413051, false},
}

func TestIsPrime(t *testing.T) {
	for _, tt := range tableTestIsPrime {
		// fmt.Printf("IsPrime(%v)\n", tt.in)
		actual := IsPrime(tt.in)
		if actual != tt.expected {
			t.Errorf("IsPrime(%v) got %v, expected %v", tt.in, actual, tt.expected)
		}
	}
}

func TestIsPrimeTrial(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	for _, ttloop := range tableTestIsPrime {
		tt := ttloop
		t.Run("Table", func(t *testing.T) {
			t.Parallel()
			actual := IsPrimeTrial(tt.in)
			if actual != tt.expected {
				t.Errorf("IsPrimeTrial(%v) got %v, expected %v", tt.in, actual, tt.expected)
			}

		})
	}
}

func TestIsMillerRabinPrime(t *testing.T) {
	for _, tt := range tableTestIsPrime {
		actual := IsMillerRabinPrime(tt.in)
		if actual != tt.expected {
			t.Errorf("IsMillerRabinPrime(%v) got %v, expected %v", tt.in, actual, tt.expected)
		}
	}
}
