package calc

// these optmized versions are probably useless

func Max(x, y int64) int64 {
	var c int64
	if x >= y {
		c = -1
	}
	return ((x ^ y) & c) ^ y
}

func MaxNaive(x, y int64) int64 {
	if x < y {
		return y
	}
	return x
}

func Min(x, y int64) int64 {
	var c int64
	if x <= y {
		c = -1
	}
	return ((x ^ y) & c) ^ y
}

func MinNaive(x, y int64) int64 {
	if x > y {
		return y
	}
	return x
}
