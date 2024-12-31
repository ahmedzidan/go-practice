package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	// 3 => 4 => 3 => 4
	uniqueList := map[int]int{}

	current := head
	for current != nil {
		uniqueList[current.Val] += 1
		current = current.Next
	}
	current2 := head
	for current2.Next != nil {
		if uniqueList[current2.Next.Val] > 1 {
			current2.Next = current2.Next.Next //1 => 1
			//current2 = current2.Next
		} else {
			current2 = current2.Next
		}
	}
	if uniqueList[head.Val] > 1 {
		newHead := head.Next
		head.Next = nil
		head = newHead
	}
	return head
}
func printHead(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		head = head.Next
	}
}
func main() {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 1}
	node3 := &ListNode{Val: 1}
	node4 := &ListNode{Val: 3}
	node5 := &ListNode{Val: 4}
	node6 := &ListNode{Val: 4}
	node7 := &ListNode{Val: 5}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = node7
	node := deleteDuplicates(node1)
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}
