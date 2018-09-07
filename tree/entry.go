package main

import (
	"fmt"
	tree "golearn/tree/entry"
)

type myTreeNode struct {
	node *tree.TreeNode
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}

	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}

func main() {
	var root tree.TreeNode

	root = tree.TreeNode{Value: 3}
	root.Left = &tree.TreeNode{}
	root.Right = &tree.TreeNode{5, nil, nil}

	root.Right.Left = new(tree.TreeNode)
	root.Left.Right = tree.CreateTreeNode(2)

	fmt.Println(root.Right.Left)

	root.Print()
	fmt.Println()
	// print(root)

	root.SetValue(5)
	// print(root)
	fmt.Println()

	var pRoot *tree.TreeNode
	pRoot.SetValue(200)
	pRoot = &root
	pRoot.SetValue(300)
	pRoot.Print()

	fmt.Println()

	root.Traverse()
	fmt.Println()

	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()

	rootCount := 0
	root.TraverseFunc(func(node *tree.TreeNode) {
		rootCount++
	})
	fmt.Println("RootCount: ", rootCount)

	// nodes := []TreeNode{
	// 	{value: 3},
	// 	{},
	// 	{6, nil, &root},
	// }
	// fmt.Println(nodes)
}
