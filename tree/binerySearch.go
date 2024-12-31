package main

import "fmt"

type Node struct {
	Key   int
	Right *Node
	Left  *Node
}

func (n *Node) insert(val int) {

	if val > n.Key {
		if n.Right != nil {
			n.Right.insert(val)
		} else {
			n.Right = &Node{Key: val}
		}

	} else if val < n.Key {
		if n.Left != nil {
			n.Right.insert(val)
		} else {
			n.Left = &Node{Key: val}
		}
	}
}
func (n *Node) search(k int) bool {
	fmt.Println(n)
	if n == nil {
		fmt.Println("eee")
		return false
	}

	if k > n.Key {
		return n.Right.search(k)
	} else if k < n.Key {
		return n.Left.search(k)
	}
	return true
}
func main() {
	tree := &Node{Key: 100}
	tree.insert(110)
	tree.insert(10)
	tree.insert(20)
	fmt.Println(tree.search(120))
	fmt.Println(tree)
}
