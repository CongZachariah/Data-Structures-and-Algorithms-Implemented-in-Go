package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print(replaceSpace("dhus ao ah s"))
}

func replaceSpace(s string) string {
	return strings.Replace(s," ","%20",-1)

}