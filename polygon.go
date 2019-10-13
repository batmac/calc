package calc

// PolygonArea computes the area of the given closed curve using
// the shoelace algorithm (points must not cross).
func PolygonArea(list []Point) float64 {
	var area float64
	length := len(list)
	if length < 3 {
		panic("a polygon must have at least 3 Points.")
	}

	for i, j := length-1, 0; j < length; i, j = j, j+1 {
		area += list[i].x*list[j].y - list[i].y*list[j].x
	}

	if area < 0 {
		area = -area
	}

	return area / 2
}
