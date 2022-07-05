package stringschallenge

type substring struct {
	left  int
	right int
}

func (ss substring) length() int {
	return ss.right - ss.left
}

// LongestSubstringWithoutDuplication takes a string and returns the longest substring without repeating characters
/*
Example 1:
Input: "clementisacap"
Output: "mentisac"
*/
func (stringsChallenge) LongestSubstringWithoutDuplication(s string) string {
	longest := substring{0, 1}
	m := make(map[rune]int)

	startIdx := 0
	for i, el := range s {

		if _, ok := m[el]; ok {
			if m[el] >= startIdx {
				startIdx = m[el] + 1
			}
		}

		if longest.length() < i-startIdx+1 {
			longest = substring{startIdx, i + 1}
		}

		m[el] = i
	}

	return s[longest.left:longest.right]
}

