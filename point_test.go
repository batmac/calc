package calc

import "testing"

var tableTestHypot = []struct {
	in       Point
	expected float64
}{
	{Point{0, 0}, 0},
	{Point{3, 4}, 5},
	{Point{-18, -24}, 30},
	{Point{-3, -4}, 5},
}

func TestHypot(t *testing.T) {
	for _, tt := range tableTestHypot {
		actual := tt.in.Hypot()
		if actual != tt.expected {
			t.Errorf("Hypot(%v) got %v, expected %v", tt.in, actual, tt.expected)
		}
	}
}
