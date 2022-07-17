package stringschallenge

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStringsChallenge_IsPalindrome(t *testing.T) {
	x := 121
	res := NewStringsChallenge().IsPalindrome(x)

	require.Equal(t, true, res)
}
