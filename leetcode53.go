package main

import (
	"fmt"
)

func main() {
	data  := []int{2}
	fmt.Print(maxSubArray( data ))
}

func maxSubArray(nums []int) int {
	max := -2147483647
	var sum int
	for l := 0 ; l<len(nums) ; l ++ {
		for r := l ; r<len(nums) ; r ++ {
			for t := l ; t<=r ; t ++{
				sum+=nums[t]
			}
			if sum > max {
				max = sum
			}
			sum = 0
		}
	}
	return max
}
