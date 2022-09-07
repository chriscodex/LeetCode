package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	s := ""
	n := strconv.Itoa(x)
	for i := len(n) - 1; i >= 0; i-- {
		s += string(n[i])
	}
	return s == n
}

func main() {
	b := isPalindrome(1221)
	fmt.Println(b)
	a := 33
	fmt.Println(a)
}
