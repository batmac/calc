package calc

import (
	"math"
)

// Mulm perform a*b (mod m), trying to avoid overflow
func Mulm(a, b, m uint64) uint64 {
	// fmt.Printf("   ->MulmS(%v,%v,%v)\n", a, b, m)
	if m == 0 {
		panic("m == 0")
	}
	if m == 1 {
		return 0
	}

	a %= m
	b %= m

	if a > b {
		a, b = b, a
	}
	if a == 0 {
		return 0
	}
	if a == 1 {
		return b
	}
	// fmt.Printf("   ->MulmN(a=%v,b=%v,m=%v) m+a-1=%v\n", a, b, m, m+a-1)
	if m+a-2 < m || a > (math.MaxUint64>>1) {
		panic("Mulm will overflow")
	}

	var r uint64
	for b > 0 {
		if b&1 == 1 {
			r = (a + r) % m
		}
		a = (a << 1) % m
		b >>= 1
	}
	return r
}

// Powm returns b^e (mod m) using fast modular exp
func Powm(b, e, m uint64) uint64 {
	// fmt.Printf(" ->Powm(%v,%v,%v)\n", b, e, m)
	if m == 0 {
		panic("m == 0")
	}
	if m == 1 {
		return 0
	}
	// if m-1 > math.MaxUint32 {
	// panic("c*b might overflow !")
	// }
	b %= m
	var c uint64 = 1
	for e > 0 {
		if e&1 == 1 {
			c = Mulm(c, b, m)
		}
		e >>= 1
		b = Mulm(b, b, m)
	}
	return c
}
