package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Print(maxAreaOfIsland([][]int {{1,1,0},
		                               {1,0,0},
	                                   {0,1,1}}))
}

func maxAreaOfIsland(grid [][]int) int {
	max := 0
	for i := 0; i < len(grid); i ++ {
		for j := 0; j < len(grid[0]); j ++ {
			if grid[i][j] == 1 {
				tmp := dfs1( grid , i , j )
				max=int(math.Max(float64(max),float64(tmp)))
			}
		}
	}
	return max
}


func dfs1(grid [][]int, i int, j int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == 0 {
		return 0
	}
	grid[i][j] = 0
	count := 1
	count += dfs1(grid, i+1, j)
	count += dfs1(grid, i-1, j)
	count += dfs1(grid, i, j+1)
	count += dfs1(grid, i, j-1)
	return count
}

