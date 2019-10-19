package calc

import (
	"math"
	"math/rand"
	"testing"
)

var tableTestFlp2 = []struct {
	in       uint64
	expected uint64
}{
	{0, 0},
	{1, 1},
	{2, 2},
	{3, 2},
	{math.MaxUint64, 1 << 63},
	{math.MaxUint32, 1 << 31},
	{math.MaxUint32 + 1, 1 << 32},
	{math.MaxUint16, 1 << 15},
	{math.MaxUint16 + 1, 1 << 16},
	{math.MaxUint8, 1 << 7},
	{math.MaxUint8 + 1, 1 << 8},
}

func TestFlp2(t *testing.T) {
	for _, tt := range tableTestFlp2 {
		actual := Flp2(tt.in)
		if actual != tt.expected {
			t.Errorf("Flp2(%v) got %v, expected %v", tt.in, actual, tt.expected)
		}
	}
}

func TestFlp2Naive1(t *testing.T) {
	for _, tt := range tableTestFlp2 {
		actual := Flp2Naive1(tt.in)
		if actual != tt.expected {
			t.Errorf("Flp2Naive1(%v) got %v, expected %v", tt.in, actual, tt.expected)
		}
	}
}

func BenchmarkFlp2Naive1(b *testing.B) {
	rand.Seed(_seed)
	n := uint64(rand.Int63())
	for i := 0; i < b.N; i++ {
		_ruint64 = Flp2Naive1(n)
	}
}

func BenchmarkFlp2(b *testing.B) {
	rand.Seed(_seed)
	n := uint64(rand.Int63())
	for i := 0; i < b.N; i++ {
		_ruint64 = Flp2(n)
	}
}

var tableTestClp2 = []struct {
	in       uint64
	expected uint64
}{
	{0, 0},
	{1, 1},
	{2, 2},
	{3, 4},
	{1 + 1<<62, 1 << 63},
	{2 + 1<<62, 1 << 63},
	{math.MaxUint32, 1 << 32},
	{1 + math.MaxUint32, 1 << 32},
	{2 + math.MaxUint32, 1 << 33},
	{math.MaxUint16, 1 << 16},
	{1 + math.MaxUint16, 1 << 16},
	{2 + math.MaxUint16, 1 << 17},
	{math.MaxUint8, 1 << 8},
	{1 + math.MaxUint8, 1 << 8},
	{2 + math.MaxUint8, 1 << 9},
}

func TestClp2(t *testing.T) {
	for _, tt := range tableTestClp2 {
		actual := Clp2(tt.in)
		if actual != tt.expected {
			t.Errorf("Clp2(%v) got %v, expected %v", tt.in, actual, tt.expected)
		}
	}
}
func BenchmarkClp2(b *testing.B) {
	rand.Seed(_seed)
	n := uint64(rand.Int63())
	for i := 0; i < b.N; i++ {
		_ruint64 = Clp2(n)
	}
}
