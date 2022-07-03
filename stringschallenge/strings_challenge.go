package stringschallenge

type stringsChallenge struct {
}

func NewStringsChallenge() *stringsChallenge {
	return &stringsChallenge{}
}

// CanConstruct given two string ransomNote and magazine, determine if ransomNote can be constructed from magazine
/*
Example 1:
Input: ransomNote = "a", magazine = "b"
Output: false

Example 2:
Input: ransomNote = "aa", magazine = "ab"
Output: false

Example 3:
Input: ransomNote = "aa", magazine = "aab"
Output: true
*/
func (stringsChallenge) CanConstruct(ransomNote string, magazine string) bool {
	var m = make(map[rune]int)
	for _, el := range magazine {
		m[el]++
	}

	for _, el := range ransomNote {
		if m[el] == 0 {
			return false
		}
		m[el]--
	}

	return true
}

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
