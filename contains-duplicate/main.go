// package main implements containsDuplicate function
package main

import (
	"fmt"
	"sort"
)

// Given an integer array nums, return true if
// any val appears at least twice in array,
// and return false if every element is distinct
func containsDuplicate(nums []int) bool {
	sort.Ints(nums)

	// return true if any value appears at least twice
	// counter? I could brute force and count dups b
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))
	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}
