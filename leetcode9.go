package main

import "fmt"

const MAX_VALUE=2147483647


func main(){
	var x int
	fmt.Scanf("%d", &x)
	fmt.Print( isPalindrome( x ) )
}

func isPalindrome(x int) bool {
	if x<0||x>MAX_VALUE{
		return false
	}
	x0 := x;
	var res  = 0
	for ; x>0 ;x/=10 {
		res = res*10 + x%10
	}

	if x0 == res {
		return true
	}
	return false
}
