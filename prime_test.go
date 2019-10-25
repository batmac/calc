package calc

import (
	"math/rand"
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
func TestIsPrimeCoherence(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	rand.Seed(_seed)
	for _, r := range []func() (int, func() uint64){
		func() (int, func() uint64) { return 100, func() uint64 { return uint64(rand.Int63()) } },
		func() (int, func() uint64) { return 10000, func() uint64 { return uint64(rand.Int31()) } },
	} {
		iters, randNumber := r()
		for i := 0; i < iters; i++ {
			n := uint64(randNumber())
			a := IsPrime(n)
			b := IsPrimeTrial(n)
			if a != b {
				t.Errorf("TestIsPrimeCoherence %v failed. %v!=%v", n, a, b)
			}
		}
	}
}

var tableTestPrimeNext = []struct {
	in               uint64
	nextExpected     uint64
	previousExpected uint64
}{
	{2, 3, 0},
	{3, 5, 2},
	{6, 7, 5},
	{24, 29, 23},
	{1024, 1031, 1021},
	{1 << 24, 16777259, 16777213},
	{1 << 31, 2147483659, 2147483647},
	{1 << 48, 281474976710677, 281474976710597},
	{1 << 62, 4611686018427388039, 4611686018427387847},
	// {1 << 63, 4611686018427388039},
}

func TestPrimeNext(t *testing.T) {
	for _, tt := range tableTestPrimeNext {
		actual := PrimeNext(tt.in)
		if actual != tt.nextExpected {
			t.Errorf("PrimeNext(%v) got %v, expected %v", tt.in, actual, tt.nextExpected)
		}
	}
}
func TestPrimePrevious(t *testing.T) {
	for _, tt := range tableTestPrimeNext {
		actual := PrimePrevious(tt.in)
		if actual != tt.previousExpected {
			t.Errorf("PrimePrevious(%v) got %v, expected %v", tt.in, actual, tt.previousExpected)
		}
	}
}
func BenchmarkIsPrimeMRPseudoPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_b = IsMillerRabinPrime(2152302898747)
	}
}
func BenchmarkIsPrimeMRBigPseudoPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_b = IsMillerRabinPrime(3825123056546413051)
	}
}
func BenchmarkIsPrimeTrialRandom31(b *testing.B) {
	rand.Seed(_seed)
	for i := 0; i < b.N; i++ {
		_b = IsPrimeTrial(uint64(rand.Int31()))
	}
}
func BenchmarkIsPrimeTrialRandom63(b *testing.B) {
	rand.Seed(_seed)
	for i := 0; i < b.N; i++ {
		_b = IsPrimeTrial(uint64(rand.Int63()))
	}
}
func BenchmarkIsPrimeRandom31(b *testing.B) {
	rand.Seed(_seed)
	for i := 0; i < b.N; i++ {
		_b = IsPrime(uint64(rand.Int31()))
	}
}
func BenchmarkIsPrimeRandom63(b *testing.B) {
	rand.Seed(_seed)
	for i := 0; i < b.N; i++ {
		_b = IsPrime(uint64(rand.Int63()))
	}
}
