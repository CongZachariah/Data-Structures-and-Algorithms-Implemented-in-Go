package main

import "fmt"

func main() {
	var target = 8
	var nums = []int {1,5,3,2,4,6,2}
	var res []int
	res = twoSum( nums , target)
	for i := range res {
		fmt.Print(res[i])
	}
}

func twoSum(nums []int, target int) []int {
	for i := 0 ; i<len(nums) ; i++ {
		for j := i ; j<len(nums) ;j++ {
			if nums[i] + nums[j] == target{
				res := []int{i , j}
				return res
			}
		}
	}
	return nil
}
