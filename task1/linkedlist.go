package main

import (
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (l *LinkedList) insert(data int) {
	newNode := &Node{
		Data: data,
	}

	if l.Head == nil {
		l.Head = newNode
		return
	}

	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

func (l LinkedList) readData() {
	current := l.Head

	for current != nil {
		fmt.Println(current.Data)
		current = current.Next
	}
}
func (l *LinkedList) deleteValue(data int) {
	current := l.Head // 1 , 2, 3
	if l.Head == nil {
		return
	}
	//current=1
	// currnet.next = node(2)
	// 2 ==2
	// 1.next = node(3)
	for current.Next != nil {
		if current.Next.Data == data {
			current.Next = current.Next.Next
			//return
		} else {
			current = current.Next
		}

	}

}
func hasNumber(head *Node, num int) bool {
	for head != nil {
		if head.Data == num {
			return true
		}
		head = head.Next
	}
	return false
}
func hasCycle(head *Node) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head      // 1
	fast := head.Next // 2

	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next      //2
		fast = fast.Next.Next // 2
	}
	return true
}
func mergeTwoLists(l1 *Node, l2 *Node) *Node {
	var FinalNode *Node
	var prev *Node
	for l1 != nil && l2 != nil {
		nextNode := &Node{}
		if l1.Data < l2.Data {
			nextNode = l1
			l1 = l1.Next
		} else {
			nextNode = l2
			l2 = l2.Next
		}
		if FinalNode == nil {
			FinalNode = nextNode
			prev = FinalNode
		} else {
			prev.Next = nextNode
			prev = nextNode
		}
		//FinalNode.insertNode(nextNode.Data)
	}
	for l1 != nil {
		prev.Next = l1
		prev = l1
		l1 = l1.Next
	}
	for l2 != nil {
		//FinalNode.insertNode(l2.Data)
		prev.Next = l2
		prev = l2
		l2 = l2.Next
	}
	return FinalNode
}
func (n *Node) insertNode(data int) {
	newNode := &Node{Data: data}
	if n == nil {
		n = newNode
	} else if n.Next == nil {
		n.Next = newNode
	} else {
		current := n.Next
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
		fmt.Println(current.Data)
	}
}
func main() {
	node1 := &Node{Data: 1}
	node2 := &Node{Data: 4}
	node3 := &Node{Data: 5}
	node1.Next = node2
	node2.Next = node3
	node4 := &Node{Data: 1}
	node5 := &Node{Data: 2}
	node6 := &Node{Data: 3}
	node4.Next = node5
	node5.Next = node6
	node := mergeTwoLists(node1, node4)
	for node != nil {
		fmt.Println(node.Data)
		node = node.Next
	}
}
