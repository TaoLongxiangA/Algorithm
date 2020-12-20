package main

import (
	"fmt"
	"strconv"
)

func main() {
	palindrome := isPalindrome(11211)
	fmt.Println(palindrome)
}

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	b := true
	var i, j int
	j = len(s) - 1
	for i = 0; i < j; i++ {
		if s[i] == s[j] {
			b = true
		} else {
			b = false
		}
		j--
	}
	return b
}
