package calc

// Flp2 rounds x down the lesser power of 2
func Flp2(x uint64) uint64 {
	x = x | (x >> 1)
	x = x | (x >> 2)
	x = x | (x >> 4)
	x = x | (x >> 8)
	x = x | (x >> 16)
	x = x | (x >> 32)
	return x - (x >> 1)
}

func Flp2Naive1(x uint64) uint64 {
	var y uint64
	for {
		y = x
		x = x & (x - 1)
		if x == 0 {
			break
		}
	}
	return y
}

// Clp2 rounds x up the next power of 2
func Clp2(x uint64) uint64 {
	x = x - 1
	x = x | (x >> 1)
	x = x | (x >> 2)
	x = x | (x >> 4)
	x = x | (x >> 8)
	x = x | (x >> 16)
	x = x | (x >> 32)
	return x + 1
}
