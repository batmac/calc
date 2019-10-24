package calc

import (
	"math"
	"math/bits"
)

// primes < 1000, without 2
var _smallPrimes = []uint64{
	3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47,
	53, 59, 61, 67, 71, 73, 79, 83, 89, 97,
	101, 103, 107, 109, 113, 127, 131, 137, 139, 149,
	151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199,
	211, 223, 227, 229, 233, 239, 241,
	251, 257, 263, 269, 271, 277, 281, 283, 293,
	307, 311, 313, 317, 331, 337, 347, 349,
	353, 359, 367, 373, 379, 383, 389, 397,
	401, 409, 419, 421, 431, 433, 439, 443, 449,
	457, 461, 463, 467, 479, 487, 491, 499,
	503, 509, 521, 523, 541, 547,
	557, 563, 569, 571, 577, 587, 593, 599,
	601, 607, 613, 617, 619, 631, 641, 643, 647,
	653, 659, 661, 673, 677, 683, 691,
	701, 709, 719, 727, 733, 739, 743,
	751, 757, 761, 769, 773, 787, 797,
	809, 811, 821, 823, 827, 829, 839,
	853, 857, 859, 863, 877, 881, 883, 887,
	907, 911, 919, 929, 937, 941, 947,
	953, 967, 971, 977, 983, 991, 997,
}

// IsPrime works with n<2^64
// fyi bpsw: http://www.trnicely.net/misc/bpsw.html
func IsPrime(n uint64) bool {
	return IsMillerRabinPrime(n)
}

// IsPrimeTrial is a trivial trial check between 2 and sqrt(n)
func IsPrimeTrial(n uint64) bool {
	if n <= 1 {
		return false
	}
	if n&1 == 0 {
		return n == 2
	}
	if n%3 == 0 {
		return n == 3
	}
	sqrtN := uint64(math.Sqrt(float64(n)))
	for i := uint64(5); i <= sqrtN; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			// fmt.Printf("%v | %v\n", i, n)
			return false
		}
	}
	return true
}

// TryVeryTrivialPrime filter [0;3] and even numbers
func TryVeryTrivialPrime(n uint64) (bool, bool) {
	sure := true
	if n <= 1 {
		// n is 0 or 1
		return sure, false
	}
	if n <= 3 {
		// n is 2 or 3
		return sure, true
	}
	if n&1 == 0 {
		// n is even
		return sure, false
	}
	return !sure, true
}

func TryTrivialPrime(n uint64) (bool, bool) {
	sure := true
	for _, v := range _smallPrimes {
		if n == v {
			return sure, true
		}
		if n%v == 0 {
			return sure, false
		}
	}
	return !sure, true
}

// IsMillerRabinPrime is a MillerRabin'ish test with n<2^64
func IsMillerRabinPrime(n uint64) bool {

	var sure, b bool
	sure, b = TryVeryTrivialPrime(n)
	if sure {
		return b
	}
	// check small prime divisors by trial
	sure, b = TryTrivialPrime(n)
	if sure {
		return b
	}

	// good-enough witnesses list for n<2^64 ( http://oeis.org/A014233 )
	bases := []uint64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37}

	// find s,D | N-1 = 2^s * D , D is odd
	Nm1 := n - 1
	s := bits.TrailingZeros64(Nm1)
	D := Nm1 >> s
WitnessLoop:
	for _, b := range bases {
		// fmt.Printf("iteration base %v , Powm(%v,%v,%v)\n", b, b, D, n)
		Rem := Powm(b, D, n)
		// fmt.Printf("b=%v, n=%v,n-1=%v,s=%v,D=%v,1<<s*D=%v, Rem=%v\n", b, n, Nm1, s, D, 1<<s*D, Rem)
		if Rem == 1 || Rem == Nm1 {
			continue WitnessLoop
		}
		for r := 1; r < s; r++ {
			// fmt.Printf("Powm(%v,2,%v)\n", Rem, n)
			Rem = Powm(Rem, 2, n)
			if Rem == Nm1 {
				// fmt.Println("Rem == Nm1")
				continue WitnessLoop
			}
		}
		// fmt.Println("n is composite")
		return false
	}
	return true
}
