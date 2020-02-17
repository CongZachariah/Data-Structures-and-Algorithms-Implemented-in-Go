package main

import "fmt"

func main() {

	var Res,data []int
	var length int
	fmt.Scanf("%d", &length)
	for i := 0; i<length ; i++ {
		fmt.Scanf("%d", data[i])
	}
    Res = plusOne(data)
    for i := 0 ; i < len(Res) ; i ++ {
	fmt.Print(Res[i])
	}
}
func plusOne(digits []int) []int {
	for i := len(digits)-1 ; i >= 0 ;i-- {
		if (digits[i]!=9) {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	return append([]int{1},digits...)
}
