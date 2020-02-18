package main

func main() {
	var nums []int
	var val,len int
	len = removeElement(nums, val);
	for i := 0; i < len; i++ {
		print(nums[i]);
	}
}
func removeElement(nums []int, val int) int {
	var j int = 0
	for i:=0;i<len(nums);i++ {
		if nums[i] != val {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}
