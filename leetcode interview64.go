package main

import "fmt"

func main() {
	fmt.Print(sumNums(10))
}

func sumNums(n int) int {
	return ( n * n + n )/2
}