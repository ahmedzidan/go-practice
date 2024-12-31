package main

type TreeNode struct {
	Val   int
	left  *TreeNode
	right *TreeNode
}

func main() {
	tree := &TreeNode{Val: 100}
	tree.insert(110)
	tree.insert(10)
	tree.insert(20)
}

func (n *TreeNode) insert(val int) {
	if val > n.Val {
		if n.right != nil {
			n.right.insert(val)
		} else {
			n.right = &TreeNode{Val: val}
		}
	} else if val < n.Val {
		if n.left != nil {
			n.left.insert(val)
		} else {
			n.left = &TreeNode{Val: val}
		}
	}

}
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.left), maxDepth(root.right))

}
