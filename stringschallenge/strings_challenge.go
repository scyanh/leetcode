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
