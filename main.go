package main

import (
	"fmt"
	"github.com/scyanh/leetcode/slidingwindow"
)

func main() {
	fmt.Println("hello")

	//str:="00110011"
	//str:="clementisacap"
	input := []int{2, 3, 1, 2, 4, 3}
	k := 7

	res := slidingwindow.NewSlidingWindowChallenge().FindSmallestSubarrayWithGivenSum(input, k)
	fmt.Println("res=", res)

}
