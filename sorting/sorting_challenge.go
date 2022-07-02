package sorting

import (
	"sort"
	"strings"
)

type sortingChallenge struct {

}

func NewSortingChallenge() *sortingChallenge {
	return &sortingChallenge{}
}

/*
ReorderLogFiles You are given an array of logs. Each log is a space-delimited string of words, where the first word is the identifier.

There are two types of logs:

Letter-logs: All words (except the identifier) consist of lowercase English letters.
Digit-logs: All words (except the identifier) consist of digits.
Reorder these logs so that:

The letter-logs come before all digit-logs.
The letter-logs are sorted lexicographically by their contents. If their contents are the same, then sort them lexicographically by their identifiers.
The digit-logs maintain their relative ordering.
Return the final order of the logs.
*/
func (s *sortingChallenge) ReorderLogFiles(logs []string) []string {
	var (
		numbers []string
		letters []string
	)
	for _, log := range logs {
		if log[len(log)-1] >= '0' && log[len(log)-1] <= '9' {
			numbers = append(numbers, log)
		} else {
			letters = append(letters, log)
		}
	}

	sort.Slice(letters, func (i, j int) bool {
		iIndex := strings.Index(letters[i], " ")
		jIndex := strings.Index(letters[j], " ")

		iLog := letters[i][iIndex+1:]
		jLog := letters[j][jIndex+1:]
		if iLog == jLog {
			return letters[i] < letters[j]
		}
		return iLog < jLog
	})

	return append(letters, numbers...)
}


