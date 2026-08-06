// Package main implements valid anagram solution
package main

import (
	"fmt"
)

// Given two strings s and t, return true if t is an anagram of s, and false otherwise
// str is arr of char
// sort alphabetically? then compare straight up
// could do 0^n brute force where I:
//   - check each letter in s against each in t
//   - everywhere theres a match, increment a counter
//   - if counter equals len of s and t then its valid
//   - have to break loop when char found
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	chars := []rune(t)

	for _, sChar := range s {
		found := false

		for j, tChar := range chars {
			if sChar == tChar {
				// Remove the matched character so it can't be used again
				chars = append(chars[:j], chars[j+1:]...)
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("aacc", "ccac"))
}
