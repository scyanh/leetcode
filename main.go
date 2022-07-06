package main

import (
	"fmt"
	"github.com/scyanh/leetcode/slidingwindow"
)

func main() {
	fmt.Println("hello")

	//str:="00110011"
	//str:="clementisacap"
	//input := []int{2, 3, 1, 2, 4, 3}
	input:="araaci"
	k := 2

	res := slidingwindow.NewSlidingWindowChallenge().LongestSubstringWithKDistinctCharacters(input, k)
	fmt.Println("res=", res)

}
