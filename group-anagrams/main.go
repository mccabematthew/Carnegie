// Package main implements group anagrams leetcode problem
package main

import (
	"fmt"
)

// counting sort, O(k) - beats general sort.Slice
func sortStrings(s string) string {
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
			result = append(result, byte('a'+i))
		}
	}
	return string(result)
}

// Given an arr of string strs, group the anagrams together
// Return answer in any order
func groupAnagrams(strs []string) [][]string {
	groupMap := make(map[string][]string)
	for _, word := range strs {
		key := sortStrings(word)
		groupMap[key] = append(groupMap[key], word)
	}
	
	groups := [][]string{}
	for _, group := range groupMap {
		groups = append(groups, group)
	}

	return groups
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(groupAnagrams([]string{""}))
	fmt.Println(groupAnagrams([]string{"a"}))
}

