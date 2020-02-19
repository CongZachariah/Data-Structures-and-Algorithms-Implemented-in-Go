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
	map1 := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		dif := target - nums[i]
		c, ok := map1[dif]
		if ok != false {
			return []int{c, i}
		}
		map1[nums[i]] = i
	}
	return nil
}
