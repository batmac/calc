package calc

import "testing"

var tableTestCheckLuhn = []struct {
	in       string
	expected bool
}{
	{"79927398713", true},
	{"  79927398713", true},
	{" 79 927398 713", true},
	{"79-927398-713", true},
	{"79-92-7398-713", true},
}

func TestCheckLuhn(t *testing.T) {
	for _, tt := range tableTestCheckLuhn {
		actual := CheckLuhn(tt.in)
		if actual != tt.expected {
			t.Errorf("CheckLuhn(%v) got %v, expected %v", tt.in, actual, tt.expected)
		}
	}
}

var tableTestGenLuhn = []struct {
	in       string
	expected int
}{
	{"7992739871", 3},
	{"79 9 2 739871", 3},
	{"79-92-739-8-71", 3},
}

func TestGenLuhn(t *testing.T) {
	for _, tt := range tableTestGenLuhn {
		actual := GenLuhn(tt.in)
		if actual != tt.expected {
			t.Errorf("GenLuhn(%v) got %v, expected %v", tt.in, actual, tt.expected)
		}
	}
}
