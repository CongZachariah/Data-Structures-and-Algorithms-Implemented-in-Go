package main

import "fmt"

func main() {
	fmt.Print(reverseLeftWords("usgfsgdosa",7))
}

func reverseLeftWords(s string, n int) string {
	return s[n:] + s[:n]
}