package main

import "fmt"

func main() {
	fmt.Print(myPow(2,-6))
}

func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	if n < 0 {
		x, n = 1 / x, -n
	}
	res := float64(1)
	for n > 0 {
		if n & 1 == 1 {  // 取n的二进制表示的最低位
			res *= x
		}
		x *= x
		n >>= 1      // 把n的二进制右移一位 == 去掉n的二进制的最低位
	}
	return res
}