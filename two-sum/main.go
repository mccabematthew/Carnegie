package main


import (
	"fmt"	
)

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target
//
// Assume that each input would have exactly one solution
// You may not use the same element twice.
// Return the answer in any order.
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target { return []int{i,j} }
		}
	}

	return nil
}

func main() {
	result0 := twoSum([]int{2,7,11,15}, 9)
	result1 := twoSum([]int{3,2,4}, 6)
	result2 := twoSum([]int{3,3}, 6)

	fmt.Println(result0)
	fmt.Println(result1)
	fmt.Println(result2)
}
// Follow-up: Can you come up with an algorithm that is less than O(n2) time complexity?

