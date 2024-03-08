package main

import "fmt"

type ListNode struct {
	val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	slowPointer, fasterPointer := head, head
	for fasterPointer != nil && fasterPointer.Next != nil {
		slowPointer = slowPointer.Next
		fasterPointer = fasterPointer.Next.Next
	}
	return slowPointer
}
func main() {
	head := &ListNode{val: 1}
	head.Next = &ListNode{val: 2}
	head.Next.Next = &ListNode{val: 3}
	head.Next.Next.Next = &ListNode{val: 4}
	head.Next.Next.Next.Next = &ListNode{val: 5}

	middle := middleNode(head)
	fmt.Println("Middle node value:", middle.val)

}
