package main

import "fmt"

func main() {
	fmt.Print(fib(100))
}

func fib(n int) int {
	data := [101]int {0,1}
	for i:= 2 ; i <= n ; i ++ {
		data[i] = (data[i-1] +data[i-2]) % (1e9 + 7)
	}
	return data[n]
}
