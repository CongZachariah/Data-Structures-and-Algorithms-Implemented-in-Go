package main

import "fmt"

func main() {
	data := []int{1,4,2,5,3,6,7,3}
	fmt.Print(maxProfit(data))
}

func maxProfit(prices []int) int {
	price := 21000000
	maxprofit := -21000000
	if len(prices) == 0{
		return 0
	}
	for i := 0 ; i < len(prices) ; i ++ {
		if prices[i] < price {
			price = prices[i]
		}
		if prices[i] - price > maxprofit {
			maxprofit = prices[i] - price
		}
	}
	return maxprofit
}