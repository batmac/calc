package calc

import (
	"math"
	"testing"
)

var tableTestAbs = []struct {
	in       int64
	expected int64
}{
	{0, 0},
	{1, 1},
	{-1, 1},
	{math.MaxInt8, math.MaxInt8},
	{math.MinInt8, math.MaxInt8 + 1},
	{math.MaxInt16, math.MaxInt16},
	{math.MinInt16, math.MaxInt16 + 1},
	{math.MaxInt32, math.MaxInt32},
	{math.MinInt32, math.MaxInt32 + 1},
	{math.MaxInt64, math.MaxInt64},
	// MinInt64 does not have a representable abs
	{math.MinInt64 + 1, math.MaxInt64},
}

func TestAbs(t *testing.T) {
	for _, tt := range tableTestAbs {
		actual := Abs(tt.in)
		if actual != tt.expected {
			t.Errorf("Abs(%v) got %v, expected %v", tt.in, actual, tt.expected)
		}
	}
}
func TestNabs(t *testing.T) {
	for _, tt := range tableTestAbs {
		actual := Nabs(tt.in)
		if actual != -tt.expected {
			t.Errorf("Nabs(%v) got %v, expected -%v", tt.in, actual, tt.expected)
		}
	}
}
