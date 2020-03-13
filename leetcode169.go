package main

import "fmt"

func main() {
	fmt.Print(majorityElement([]int {3,54,3,2,2,2,3,3,3,3,3,5,2}))
}

func majorityElement(nums []int) int {
	res := 0
	cnt := 0
	for i := 0 ; i < len(nums) ; i ++ {
		if cnt==0 {
			res = nums[i]
			cnt ++
		} else{
			if res == nums[i] {
				cnt ++
			}  else{
				cnt--
			}
		}
	}
	return res
}