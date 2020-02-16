package main

import "fmt"

const (
	MAX_VALUE=2147483647
	MIN_VALUE=-2147483648
)

func main() {
	var x int
	var s  = 0
	fmt.Scanf("%d",&x)
	for ; x!=0 ;x/=10 {
		s = s*10 + x%10
	}
	if s<MIN_VALUE||s>MAX_VALUE{
		fmt.Print(0)
	}else{
		fmt.Print(s)
	}
}
