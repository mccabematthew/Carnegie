package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	minPriceSoFar := prices[0]
	maxProfitSoFar := 0

	for _, price := range prices {
		if price < minPriceSoFar {
			minPriceSoFar = price
		} else if price-minPriceSoFar > maxProfitSoFar {
			maxProfitSoFar = price - minPriceSoFar
		}
	}

	return maxProfitSoFar
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
	fmt.Println(maxProfit([]int{3, 7, 3, 4, 1, 8}))
}
