package main

import "fmt"

func main() {

	fmt.Print(minArray([]int {3,4,5,6,7,1,2}))
}

func minArray(numbers []int) int {
	min := 0
	for i := 1; i < len(numbers); i++ {
		if numbers[i] < numbers[min] {
			min = i
			break
		}
	}
	return numbers[min]
}