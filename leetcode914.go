package main

import "fmt"

func main() {
	fmt.Print(hasGroupsSizeX([]int {1,2,1,1,1,4,3,4,4,3,3,3,2,2,2}))
}

func hasGroupsSizeX(deck []int) bool {
	map1 := make(map[int]int)
	for i := 0 ; i < len(deck) ; i ++ {
		map1[deck[i]] ++
	}
	g := 0
	for i := 1 ; i <= len(map1) ; i ++ {
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