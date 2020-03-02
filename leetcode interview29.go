package main

import "fmt"

func main() {
	data := [][]int {{1,2,3},{4,5,6},{7,8,9}}
	res :=  spiralOrder(data)
	fmt.Print(res)
}

func spiralOrder(matrix [][]int) []int { // 本函数用于顺时针遍历一个二维数组，并把遍历过程用一维数组返回
	if len ( matrix ) == 0{
		return nil  // 对于空二维数组的特判，直接返回空即可。
	}
	step := 0
	size := len( matrix ) * len( matrix[ 0 ] )
	// step 、size代表当前已经走过的步数和要走的总步数。
	
	top, bottom, right, left := 0, len(matrix)-1 ,len(matrix[0])-1 ,0
	// 定义横向的起、终点和纵向的起、终点
	nums := make([]int,size)
	for step < size {
		// 本for循环模拟从左走到右
		for i := left ; i <= right && step < size ; i ++ {
			nums[step] = matrix[top][i]  // 记录走过的路径，下同
			step ++ 
		}
		top ++
		// 刚刚从左到右走完，因此上边界向下移动一格，下同。
		// 同理，从上到下
		for i := top ; i <= bottom && step < size ; i ++ {
			nums[step] = matrix[i][right]
			step ++
		}
		right --
		// 同理，从右到左
		for i := right ; i >= left && step < size ; i -- {
			nums[step] = matrix[bottom][i]
			step ++
		}
		bottom --
		// 同理，从下到上
		for i := bottom ; i >= top && step < size ; i -- {
			nums[step] = matrix[i][left]
			step ++
		}
		left ++
	}
	return nums  // 返回路径记录数组 nums
}
