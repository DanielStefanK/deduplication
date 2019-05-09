package calculator

import "unicode/utf8"

func ComputeDistance(a, b string) int {
	if len(a) == 0 {
		return utf8.RuneCountInString(b)
	}

	if len(b) == 0 {
		return utf8.RuneCountInString(a)
	}

	if a == b {
		return 0
	}

	s1 := []rune(a)
	s2 := []rune(b)

	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}

	lenS1 := len(s1)
	lenS2 := len(s2)

	x := make([]int, lenS1+1)
	for i := 0; i < len(x); i++ {
		x[i] = i
	}

	_ = x[lenS1]

	for i := 1; i <= lenS2; i++ {
		prev := i
		var current int
		for j := 1; j <= lenS1; j++ {
			if s2[i-1] == s1[j-1] {
				current = x[j-1] // match
			} else {
				current = min(min(x[j-1]+1, prev+1), x[j]+1)
			}
			x[j-1] = prev
			prev = current
		}
		x[lenS1] = prev
	}
	return x[lenS1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
