package main

import (
	"fmt"
	"github.com/scyanh/leetcode/slidingwindow"
)

func main() {
	fmt.Println("hello")

	//str:="00110011"
	//str:="clementisacap"
	input:=[]int{1, 3, 2, 6, -1, 4, 1, 8, 2}
	k:=5

	res := slidingwindow.NewSlidingWindowChallenge().FindAverageOfSubarrayOfSizeK(input, k)
	fmt.Println("res=", res)

}
