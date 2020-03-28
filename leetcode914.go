package main

import "fmt"

func main() {
	fmt.Print(hasGroupsSizeX([]int {0,0,0,0,0}))
}

func hasGroupsSizeX(deck []int) bool {
	map1 := make(map[int]int)
	for i := 0 ; i < len(deck) ; i ++ {
		map1[deck[i]] ++
	}
	g := 0
	for i := 0 ; i <= len(map1) ; i ++ {
		fmt.Print(map1[i])
		g = gcd(g,map1[i])
		if g == 1 {
			return false
		}
	}
	return g >= 2
}

func gcd(a,b int) int {
	if b == 0 {
		return a
	}
	return gcd( b,a % b )
}