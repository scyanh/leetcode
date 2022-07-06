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
	input:=[]string{"A","B","C","A","C"}
	//k := 2

	res := slidingwindow.NewSlidingWindowChallenge().MaxFruitsIntoBasket(input)
	fmt.Println("res=", res)

}
