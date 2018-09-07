package tree

import (
	"fmt"
)

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

func (node TreeNode) Print() { //receiver is TreeNode
	fmt.Print(node.Value, " ")
}

func Print(node TreeNode) {
	fmt.Println(node.Value)
}

func (node *TreeNode) SetValue(value int) { // receiver is *TreeNode
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored.")
		return
	}
	node.Value = value
}

func (node *TreeNode) Traverse() {
	// if node == nil {
	// 	return
	// }
	// node.Left.Traverse()
	// node.Print()
	// node.Right.Traverse()
	node.TraverseFunc(func(tree *TreeNode) {
		tree.Print()
	})
	fmt.Println("------------------")
}

func (node *TreeNode) TraverseFunc(f func(*TreeNode)) {
	if node == nil {
		return
	}
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}

func CreateTreeNode(value int) *TreeNode {
	return &TreeNode{Value: value}
}
