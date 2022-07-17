package main

import (
	"fmt"
	"github.com/scyanh/leetcode/stringschallenge"
)

func main() {
	fmt.Println("hello")

	//str:="00110011"
	//str:="clementisacap"
	//input := []int{2, 3, 1, 2, 4, 3}
	//input:=[]string{"A","B","C","A","C"}
	//k := 2

	x := 121

	res := stringschallenge.NewStringsChallenge().IsPalindrome(x)
	fmt.Println("res=", res)

}
