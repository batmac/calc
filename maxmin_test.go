package calc

import (
	"math"
	"math/rand"
	"testing"
)

var tableTestMax = []struct {
	x, y     int64
	expected int64
}{
	{1, 2, 2},
	{2, 1, 2},
	{1, 1, 1},
	{-1, -2, -1},
	{-2, -1, -1},
	{-1, -1, -1},
	{0, 0, 0},
	{-1, 2, 2},
	{2, -1, 2},
	{1, -1, 1},
	{-1, 1, 1},
	// {math.MaxInt64, math.MinInt64 + 1, math.MaxInt64},
}

func TestMax(t *testing.T) {
	for _, tt := range tableTestMax {
		actual := Max(tt.x, tt.y)
		if actual != tt.expected {
			t.Errorf("Max(%v,%v) got %v, expected %v", tt.x, tt.y, actual, tt.expected)
		}
	}
}

var tableTestMin = []struct {
	x, y     int64
	expected int64
}{
	{1, 2, 1},
	{2, 1, 1},
	{1, 1, 1},
	{-1, -2, -2},
	{-2, -1, -2},
	{-1, -1, -1},
	{0, 0, 0},
	{-1, 2, -1},
	{2, -1, -1},
	{1, -1, -1},
	{-1, 1, -1},
	{math.MaxInt64, math.MinInt64, math.MinInt64},
}

func TestMin(t *testing.T) {
	for _, tt := range tableTestMin {
		actual := Min(tt.x, tt.y)
		if actual != tt.expected {
			t.Errorf("Min(%v,%v) got %v, expected %v", tt.x, tt.y, actual, tt.expected)
		}
	}
}

func BenchmarkMax(b *testing.B) {
	rand.Seed(_seed)
	for i := 0; i < b.N; i++ {
		_rint64 = Max(rand.Int63(), rand.Int63())
	}
}
func BenchmarkMaxNaive(b *testing.B) {
	rand.Seed(_seed)
	for i := 0; i < b.N; i++ {
		_rint64 = MaxNaive(rand.Int63(), rand.Int63())
	}
}

func BenchmarkMin(b *testing.B) {
	rand.Seed(_seed)
	for i := 0; i < b.N; i++ {
		_rint64 = Min(rand.Int63(), rand.Int63())
	}
}
func BenchmarkMinNaive(b *testing.B) {
	rand.Seed(_seed)
	for i := 0; i < b.N; i++ {
		_rint64 = MinNaive(rand.Int63(), rand.Int63())
	}
}
