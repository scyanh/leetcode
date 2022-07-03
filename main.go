package main

import (
	"fmt"
	"github.com/scyanh/leetcode/stringschallenge"
)

func main() {
	fmt.Println("hello")

	//str:="00110011"
	str:="clementisacap"

	res := stringschallenge.NewStringsChallenge().LongestSubstringWithoutDuplication(str)
	fmt.Println("res=", res)

}
