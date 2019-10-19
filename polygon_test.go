package calc

import "testing"

var tableTestPolygonArea = []struct {
	in       []Point
	expected float64
}{
	// check types
	{[]Point{{float64(0), float64(0)}, {float64(1), float64(0)}, {float64(1), float64(1)}, {0, 1}}, float64(1)},
	// 3 points
	{[]Point{{1, 0}, {1, 1}, {0, 1}}, 0.5},
	// rotating points from a square
	{[]Point{{0, 0}, {1, 0}, {1, 1}, {0, 1}}, 1},
	{[]Point{{1, 0}, {1, 1}, {0, 1}, {0, 0}}, 1},
	{[]Point{{1, 1}, {0, 1}, {0, 0}, {1, 0}}, 1},
	{[]Point{{0, 1}, {0, 0}, {1, 0}, {1, 1}}, 1},
	// random leveled rectangle
	{[]Point{{5, 5}, {-1, 5}, {-1, -2}, {5, -2}}, 42},
	// random polygon
	{[]Point{{5, -1}, {2, 4}, {-1, 2}, {2, 0}}, 12},
	// points in reverse order
	{[]Point{{2, 0}, {-1, 2}, {2, 4}, {5, -1}}, 12},
}

func TestPolygonArea(t *testing.T) {
	for _, tt := range tableTestPolygonArea {
		actual := PolygonArea(tt.in)
		if actual != tt.expected {
			t.Errorf("PolygonArea(%v) got %v, expected %v", tt.in, actual, tt.expected)
		}
	}
}
