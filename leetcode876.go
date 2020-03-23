package main

type ListNode struct {
	     Val int
	     Next *ListNode
	 }

func main() {
	var data *ListNode
	middleNode(data)
}

func middleNode(head *ListNode) *ListNode {
	var slow = head
	var fast = head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}