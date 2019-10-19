package calc

import (
	"math"
	"math/rand"
	"testing"
)

var tableTestUgcd = []struct {
	a, b     uint64
	expected uint64
}{
	{0, 0, 0},
	{0, 1, 1},
	{1, 0, 1},
	{1, 1, 1},
	{math.MaxUint64, 140071481932319848, 1},
	{math.MaxUint64, math.MaxInt32, 1},
	{math.MaxUint64, math.MaxInt64, 1},
	{4 * math.MaxUint32, 4 * math.MaxInt32, 4},
	{math.MaxUint64, math.MaxUint64, math.MaxUint64},
	{math.MaxUint64, math.MaxUint64 - 1, 1},
	{math.MaxUint64 >> 1, math.MaxUint64 / 2, math.MaxUint64 >> 1},
	{(math.MaxUint64 - 1) >> 1, (math.MaxUint64 - 1) / 2, (math.MaxUint64 - 1) >> 1},
	{18446744073709551614, 18446744073709551614 / 2, 18446744073709551615 >> 1},
}

func TestUgcdEuclid(t *testing.T) {
	for _, tt := range tableTestUgcd {
		actual := UgcdEuclid(tt.a, tt.b)
		if actual != tt.expected {
			t.Errorf("Ugcd(%v,%v) got %v, expected %v", tt.a, tt.b, actual, tt.expected)
		}
	}
}

func TestUgcdBinary(t *testing.T) {
	for _, tt := range tableTestUgcd {
		actual := UgcdBinary(tt.a, tt.b)
		if actual != tt.expected {
			t.Errorf("UgcdBinary(%v,%v) got %v, expected %v", tt.a, tt.b, actual, tt.expected)
		}
	}
}

func TestUgcdConsistence(t *testing.T) {
	rand.Seed(_seed)
	for i := 0; i < 100000; i++ {
		a, b := uint64(rand.Int63()), uint64(rand.Int63())
		u := UgcdEuclid(a, b)
		v := UgcdBinary(a, b)
		if u != v {
			t.Errorf("a=%v, b=%v, Ugcd=%v, UgcdBinary=%v\n", a, b, u, v)
		}
	}
}

func BenchmarkUgcdEuclid(b *testing.B) {
	rand.Seed(_seed)
	u, v := uint64(rand.Int63()), uint64(rand.Int63())
	for i := 0; i < b.N; i++ {
		_r = UgcdEuclid(u, v)
	}
}
func BenchmarkUgcdBinary(b *testing.B) {
	rand.Seed(_seed)
	u, v := uint64(rand.Int63()), uint64(rand.Int63())
	for i := 0; i < b.N; i++ {
		_r = UgcdBinary(u, v)
	}
}

var tableTestLcm = []struct {
	a, b     int64
	expected int64
}{
	{0, 0, 0},
	{0, 1, 0},
	{1, 0, 0},
	{1, 1, 1},
	{10, 15, 30},
}

func TestLcm(t *testing.T) {
	for _, tt := range tableTestLcm {
		actual := Lcm(tt.a, tt.b)
		if actual != tt.expected {
			t.Errorf("Lcm(%v,%v) got %v, expected %v", tt.a, tt.b, actual, tt.expected)
		}
	}
}
