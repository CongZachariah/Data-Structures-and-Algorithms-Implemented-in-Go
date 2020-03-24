package main

import "fmt"

func main() {
	fmt.Print(massage([]int {1,5,23,2,1,5,34,2,5,6,6,4,1}))
}

func massage(nums []int) int {
	if len(nums) == 0 {
		return 0
	}else if len(nums) == 1 {
		return nums[0]
	}
	var dp = make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = maxA(nums[0] , nums[1])
	for i := 2 ; i < len(nums) ; i ++ {
		dp[i] = maxA(dp[i-2]+nums[i],dp[i-1])
	}
	return dp[len(nums)-1]
}

func maxA(a int,b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}