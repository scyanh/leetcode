package sorting

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSortingChallenge_ReorderLogFiles(t *testing.T) {
	sortingChallenge:=NewSortingChallenge()

	arr := []string{"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero"}
	res := sortingChallenge.ReorderLogFiles(arr)

	arr2 := []string{"a1 9 2 3 1","g1 act car","zo4 4 7","ab1 off key dog","a8 act zoo"}
	res2 := sortingChallenge.ReorderLogFiles(arr2)

	expected := []string{"let1 art can", "let3 art zero", "let2 own kit dig", "dig1 8 1 5 1", "dig2 3 6"}
	expected2 := []string{"g1 act car","a8 act zoo","ab1 off key dog","a1 9 2 3 1","zo4 4 7"}

	require.Equal(t, expected, res)
	require.Equal(t, expected2, res2)
}
