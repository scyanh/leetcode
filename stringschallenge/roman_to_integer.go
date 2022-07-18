package stringschallenge

func (stringsChallenge) RomanToInteger(str string) int {
	var res int

	memo := make(map[rune]int)
	memo['I'] = 1
	memo['V'] = 5
	memo['X'] = 10
	memo['L'] = 50
	memo['C'] = 100
	memo['D'] = 500
	memo['M'] = 1000

	i := 0
	for i < len(str) {
		if i < len(str)-1 && memo[rune(str[i])] < memo[rune(str[i+1])] {
			res += subsEl(memo, rune(str[i]), rune(str[i+1]))
			i += 2
		} else {
			res += memo[rune(str[i])]
			i++
		}
	}
	return res
}

func subsEl(memo map[rune]int, a, b rune) int {
	return memo[b] - memo[a]
}
