package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Print(compressString("aaaabbxxxx"))

}

func compressString(S string) string {
	if len(S) == 0{
		return S
	}
	ans := ""
	cnt := 1
	ch := S[0]
	for i := 1; i < len(S); i++ {
		if ch == S[i] {
			cnt++
		} else{
			ans += string(ch) + strconv.Itoa(cnt)
			ch = S[i]
			cnt = 1
		}
	}
	ans += string(ch) + strconv.Itoa(cnt)
	if len(ans) >= len(S) {
		return S
	} else {
		return ans
	}
}