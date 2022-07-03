package main

import (
	"fmt"
	"github.com/scyanh/leetcode/stringschallenge"
)

func main() {
	fmt.Println("hello")

	//str:="00110011"
	str:="01"

	res := stringschallenge.NewStringsChallenge().CountBinarySubStrings(str)
	fmt.Println("res=", res)

}
