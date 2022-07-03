package stringschallenge

type stringsChallenge struct {
}

func NewStringsChallenge() *stringsChallenge {
	return &stringsChallenge{}
}

// CountBinarySubStrings
/*
Given a binary string s, return the number of non-empty substrings that have the same number of 0's and 1's,
and all the 0's and all the 1's in these substrings are grouped consecutively.

Substrings that occur multiple times are counted the number of times they occur.

Input: s = "00110011"
Output: 6
Explanation: There are 6 substrings that have equal number of consecutive 1's and 0's: "0011", "01", "1100", "10", "0011", and "01".
Notice that some of these substrings repeat and are counted the number of times they occur.
Also, "00110011" is not a valid substring because all the 0's (and 1's) are not grouped together.

*/
func (sc stringsChallenge) CountBinarySubStrings(s string) int {
	current := 1
	precedent := 0
	result := 0

	for i := 1; i < len(s); i++ {
		if s[i-1] == s[i] {
			current++
		} else {
			result += sc.min(precedent, current)
			precedent, current = current, 1
		}
	}
	return result + sc.min(precedent, current)
}

func (stringsChallenge) min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
