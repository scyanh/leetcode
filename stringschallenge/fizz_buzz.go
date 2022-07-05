package stringschallenge

import "fmt"

// FizzBuzz Given a number n, for each integer i in the range from 1 to n inclusive, print one value per line as follows:
/*
if i is a multiple of both 3 and 5, print FizzBuss.
if i is a multiple of 3, print Fizz.
if i is a multiple of 5, print Buzz.
if i is not a multiple of 3 or 5, print the value of i.
*/
func (stringsChallenge) FizzBuzz(n int) []string {
	var res []string
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			res = append(res, "FizzBuzz")
		} else if i%3 == 0 {
			res = append(res, "Fizz")
		} else if i%5 == 0 {
			res = append(res, "Buzz")
		} else {
			res = append(res, fmt.Sprintf("%d", i))
		}
	}
	return res
}
