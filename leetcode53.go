package main

import (
	"fmt"
)

func main() {
	data  := []int{2}
	fmt.Print(maxSubArray( data ))
}

func maxSubArray(nums []int) int {
	if len(nums)<2{
		return nums[0]
	}
	var temp int
	res := -2147483647
	for i:= 0 ;i<len(nums);i++{
		if nums[i]>temp+nums[i]{
			temp = nums[i]
		}else {
			temp = temp+nums[i]
		}
		if temp>res{
			res = temp
		}
	}
	return res
}
