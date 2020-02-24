package main

import "fmt"

func main() {
	data := [][]int { {3,2,3,2}, {1,6,1,4},{3,5,4,3}}
	fmt.Print(maxValue(data))
}

func maxValue(grid [][]int) int {
	if len(grid) == 0{
		return 0
	}
	dp := make([][]int, len(grid)+1)
	for i := 0; i < len(dp); i ++{
		dp[i] = make([]int, len(grid[0])+1)
	}
	dp[0][0] = grid[0][0]
	for i := 0; i < len(grid); i ++ {
		for j:=0; j<len(grid[0]);j++{
			if dp[i][j+1] > dp[i+1][j]{
				dp[i+1][j+1] = dp[i][j+1] + grid[i][j]
			} else {
				dp[i+1][j+1] = dp[i+1][j] + grid[i][j]
			}

		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
	/*m := len(grid)
	n := len(grid[0])
	if m == 0 {
		return 0
	}
	var res = make([][]int, m+1)
	res[0][0] = grid[0][0]
	for i := 0 ; i < len(grid) ; i++{
		res[i] = make([]int, len(grid[0])+1)
	}
	for i := 0 ; i < m ; i ++ {
		for j := 0 ; j < n ;j ++ {
			if res[i+1][j] > res[i][j+1] {
				res[i+1][j+1] = grid[i][j] + res[i+1][j]
			} else {
				res[i+1][j+1] = grid[i][j] + res[i][j+1]
			}
		}
	}
	return res[m-1][n-1]*/
}