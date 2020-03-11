package main

import "fmt"

func main(){
	fmt.Print(constructArr([]int {4,2,6,4,2,7,7,8,6,5}))
}

func constructArr(a []int) []int {
	var b = make([]int, len(a))
	var left = 1
	for i := range a {
		b[i] = left
		left *= a[i]
	}
	var right = 1
	for j := len(a) - 1; j >= 0; j-- {
		b[j] *= right
		right *= a[j]
	}
	return b
}
