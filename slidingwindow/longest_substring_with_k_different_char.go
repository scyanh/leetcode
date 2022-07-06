package slidingwindow

// LongestSubstringWithKDistinctCharacters finds the longest substring with k distinct characters
// Example 1:
// Input: s = "araaci", k = 2
// Output: 4
// Time Complexity: O(n+n) = O(n)
// Space Complexity: O(k) = O(n)
func (slidingWindowChallenge) LongestSubstringWithKDistinctCharacters(s string, k int) int {
	n := len(s) // length of string
	if n == 0 {
		return 0 // if string is empty, return 0
	}

	left := 0 // left pointer
	res := 0 // result
	count := 0 // count of distinct characters
	m := make(map[rune]int) // map of characters and their counts

	for i := 0; i < n; i++ { // iterate through string
		r := rune(s[i]) // convert character to rune
		if _, ok := m[r]; ok { // if character is already in map
			m[r]++ // increment count
		} else {
			m[r] = 1 // else add new character to map
			count++ // increment count
		}

		for count > k { // if count is greater than k
			r := rune(s[left]) // convert character to rune
			if _, ok := m[r]; ok { // if character is already in map
				m[r]-- // decrement count
				if m[r] == 0 {
					count-- // decrement count
				}
			}
			left++ // move left pointer
		}

		res = max(res, i+1-left) // update result
	}

	return res // return result
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
