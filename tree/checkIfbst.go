package main

type BSTNode struct {
	Key   int
	Right *Node
	Left  *Node
}

func isBST(tree BSTNode) bool {
	base := tree.Key
	if base < tree.Left.Key {
		return false
	} else if base > tree.Right.Key {
		return false
	}
	return true
}
