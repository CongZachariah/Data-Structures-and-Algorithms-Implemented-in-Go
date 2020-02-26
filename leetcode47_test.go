package gotest

import "testing"

func TestGetValue(t *testing.T) {
	value := maxValue([][]int { {3,2,3,2}, {1,6,1,4},{3,5,4,3}})
	if value != 23 {
		t.Error("测试失败")
	}
}
