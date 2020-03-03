package main

import "fmt"

func main() {
	fmt.Print(movingCount(16,8,4))
}

func movingCount(m int, n int, k int) int {
	// 首先初始化一个二维数组，用于标记二维矩阵的每个格子是否已经走过了。
	visit := make([][]bool, m + 1)
	for i := 0; i < len(visit); i++ {
		visit[i] = make([]bool, n + 1)
	}
	//进行深搜。
	return dfs(0, 0, m, n, k, visit)
}

// 本函数用于对二维矩阵进行深度优先搜索。
func dfs(x, y, m, n, k int, visit [][]bool) int {
	if x >= m || y >= n || x < 0 || y < 0 {
		return 0 // 如果这个地方越界了，就返回
	}
	if visit[x][y] {
		return 0 // 如果这个地方已经被访问过了，就返回
	}
	if splitData(x, y) > k {
		return 0 // 如果这个地方的坐标拆分之和不满足题目要求，就返回
	}
	visit[x][y] = true //标记自己已经走过这个格子了
	sum := 1
	// 深度优先搜索法，按四个方向进行搜索，并累积满足要求的格子数量。
	// sum += dfs(x - 1, y, m, n, k, visit)
	sum += dfs(x + 1, y, m, n, k, visit)
	//sum += dfs(x, y - 1, m, n, k, visit)
	sum += dfs(x, y + 1, m, n, k, visit)
	return sum
}

// 本函数用于对坐标的值进行按位拆分，并求和
func splitData(m int, n int) int {
	sum := 0
	for ; m > 0 ; m /= 10 {
		sum += m%10
	}
	for ; n > 0 ; n /= 10 {
		sum += n%10
	}
	return sum
}