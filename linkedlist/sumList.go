package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	finalList := &ListNode{}
	current := finalList
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
		carry = sum / 10
	}

	return finalList.Next
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var finalNode *ListNode
	//currentNode := finalNode
	reminder := 0
	current := finalNode
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + reminder
		if sum >= 10 {
			sum -= 10
			reminder = 1
		} else {
			reminder = 0
		}
		newNode := &ListNode{Val: sum}
		if finalNode == nil {
			finalNode = newNode
			current = newNode
		} else {
			current.Next = newNode
			current = newNode
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		sum := reminder + l1.Val
		if sum > 10 {
			sum -= 10
			reminder = 1
		} else {
			reminder = 0
		}
		newNode := &ListNode{
			Val: sum,
		}
		current.Next = newNode
		current = newNode
		l1 = l1.Next
	}

	for l2 != nil {
		sum := reminder + l2.Val
		if sum > 10 {
			sum -= 10
			reminder = 1
		} else {
			reminder = 0
		}
		newNode := &ListNode{
			Val: sum,
		}
		current.Next = newNode
		current = newNode
		l2 = l2.Next
	}
	return finalNode
}

func main() {
	fmt.Println(7 / 10)
	node1 := &ListNode{Val: 2}
	node2 := &ListNode{Val: 4}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 3}

	node5 := &ListNode{Val: 5}
	node6 := &ListNode{Val: 6}
	node7 := &ListNode{Val: 4}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node5.Next = node6
	node6.Next = node7
	node := addTwoNumbers2(node1, node5)
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}
