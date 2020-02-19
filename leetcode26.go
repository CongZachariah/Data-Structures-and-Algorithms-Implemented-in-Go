package main

import "fmt"

func main() {
	data := []int {1,1,2}
	res := removeDuplicates(data)
	fmt.Print(res )
}

func removeDuplicates(nums []int) int {
	for i := 0; i+1 < len(nums); {
		if nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
		i++
		}
	}
	return len(nums)
}