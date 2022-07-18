package stringschallenge

import "strconv"

func (stringsChallenge) IsPalindrome(x int) bool {
	str := strconv.Itoa(x)

	size := len(str)
	half := size / 2

	for i := 0; i < half; i++ {
		if str[i] != str[size-i-1] {
			return false
		}
	}

	return true
}
