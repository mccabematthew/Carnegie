// Package main implements LeetCode valid palindrome solution
package main

import (
	"fmt"
)

func validPalindrome(s string) bool {
	left, right := 0, len(s)-1

	isAlnum := func(b byte) bool {
		return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9')
	}

	toLower := func(b byte) byte {
		if b >= 'A' && b <= 'Z' {
			return b + ('a' - 'A')
		}
		return b
	}

	for left < right {
		for left < right && !isAlnum(s[left]) {
			left++
		}
		for left < right && !isAlnum(s[right]) {
			right--
		}

		if toLower(s[left]) != toLower(s[right]) {
			return false
		}

		left++
		right--
	}

	return true
}

func main() {
	fmt.Println(validPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(validPalindrome("race a car"))
	fmt.Println(validPalindrome(" "))
}

