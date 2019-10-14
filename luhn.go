package calc

import "strings"

func CheckLuhn(s string) bool {
	l := len(s) - 1
	return int(s[l]-'0') == GenLuhn(s[:l])
}

// GenLuhn returns the Luhn number needed to pass the checksum
func GenLuhn(s string) int {
	var sum int
	sumTable := [...]int{0, 2, 4, 6, 8, 1, 3, 5, 7, 9}

	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "-", "")

	odd := len(s) & 1
	for i, c := range s {
		if c < '0' || c > '9' {
			return -1
		}

		if i&1 == odd {
			sum += int(c - '0')
		} else {
			sum += sumTable[c-'0']
		}
	}
	return 10 - sum%10
}
