package main

import "fmt"

func main() {
	fmt.Print(surfaceArea([][]int{{2,3,1},{7,6,4},{1,0,1}}))

}

func surfaceArea(grid [][]int) int {
	if len(grid) == 0{
		return 0
	}

	count := 0
	sum := 0

	for i:=0;i<len(grid);i++ {
		for j:=0;j<len(grid);j++ {
			if grid[i][j] == 0{
				continue
			}
			sum += grid[i][j]
			count += grid[i][j] - 1
			if j-1 >=0 {
				count += Min(grid[i][j-1],grid[i][j])
			}
			if i-1 >=0 {
				count += Min(grid[i-1][j],grid[i][j])
			}
			// fmt.Println(i,j,count,sum)
		}
	}
	return sum*6 - count*2
}

func Min(a,b int) int {
	if a<b {
		return a
	}
	return b
}