package calc

// Gcd returns the (a & b are signed) gcd
func Gcd(a, b int64) int64 {
	a, b = Abs(a), Abs(b)
	return int64(UgcdBinary(uint64(a), uint64(b)))
}

// Ugcd returns the (a & b are unsigned) pgcd
func UgcdEuclid(a, b uint64) uint64 {
	if a < b {
		a, b = b, a
	}
	if b == 0 {
		return a
	}
	return UgcdEuclid(b, a%b)
}

// UgcdBinary is a non-recursive binary (stein) algorithm
func UgcdBinary(a, b uint64) uint64 {
	// special cases
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}

	var d int

	// remove common 2^d as divisor
	for (a|b)&1 == 0 {
		d++
		a, b = a>>1, b>>1
	}

	// now we know one of a or b is not divisible by 2

	// remove remaining 2^M (if any) in a
	for a&1 == 0 {
		a >>= 1
	}

	for {
		// remove remaining 2^N (if any) in b
		for b&1 == 0 {
			b >>= 1
		}
		if a > b {
			a, b = b, a
		}
		b -= a
		if b == 0 {
			break
		}
	}

	return a << d
}

// Lcm returns the lcm of a and b
//  quick&dirty impl considering innt64 only (not uint64)
func Lcm(a, b int64) int64 {
	if a == 0 || b == 0 {
		return 0
	}
	return Abs(a / Gcd(a, b) * b)
}
