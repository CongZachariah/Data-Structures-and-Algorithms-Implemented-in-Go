package main

import "fmt"

func main() {
	var str = "s  "
	var res int
	res = lengthOfLastWord(str)
	fmt.Print(res)
}

func lengthOfLastWord(s string) int {
	var count int
	for i := len(s)-1 ; i >= 0 ; i -- {
		if s[i] != ' ' {
			count++
		} else if count>0 {
			return count
		}
	}
	return count
}
