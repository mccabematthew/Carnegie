// Package main implements group anagrams leetcode problem
package main

import (
	"fmt"
)

func sortStrings(s string) string{
	// immutable by length, mutable by value (zero initialized)
	var counts [26]int
	for _, c := range s {
		// char arithmetic using rune literals, 'a' evals to 97 (unicode)
		// 	- say c = 'e' (which is 101), then 'e' - 'a' == 4
		// 	- arr indices start at 0, making 4 the 5th index. e is 5th letter in alphabet
		//	- tally is incremented at each index, 'e' goes at 4, 'a' goes at 0
		//	- so counts for the word 'eat' looks like [1,0,0,0,1,0,...,1,...]
		counts[c-'a']++
	}
	result := make([]byte, 0, len(s)) // byte slice, 0 is start len, len(s) is capacity (prealloc space)
	for i, count := range counts {
		for j := 0; j < count; j++ {
			result = append(result, byte('a'+i)
		}
	}
	return string(result)
}

// Given an arr of string strs, group the anagrams together
// Return answer in any order
func groupAnagrams(strs []string) [][]string {
	// find anagrams
	// build arr of string arrs around anagrams as we find them
	groups := [][]string{}

	// sort alphabetically may be best course in this situation
	// We can also check every char against every char but that requires work around for dups
	

	


	return groups
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(groupAnagrams([]string{""}))
	fmt.Println(groupAnagrams([]string{"a"}))
}
