package tree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func initTree(root *TreeNode) *TreeNode {
	l1 := TreeNode{}
	r2 := TreeNode{}
	root.Left = l1.NewTree(2)
	root.Right = r2.NewTree(3)
	l2 := TreeNode{}
	ll2 := TreeNode{}
	root.Left.Left = l2.NewTree(4)
	root.Left.Right = ll2.NewTree(5)
	return root
}

func (t *TreeNode) NewTree(val int) *TreeNode {
	t.Val = val
	t.Left = nil
	t.Right = nil
	return t
}

func preOrderTraversalRec(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Val)
	preOrderTraversalRec(root.Left)
	preOrderTraversalRec(root.Right)
}

func inOrderTraversalRec(root *TreeNode) {
	if root == nil {
		return
	}
	inOrderTraversalRec(root.Left)
	fmt.Printf("%d ", root.Val)
	inOrderTraversalRec(root.Right)
}

func postOrderTraversalRec(root *TreeNode) {
	if root == nil {
		return
	}
	postOrderTraversalRec(root.Left)
	postOrderTraversalRec(root.Right)
	fmt.Printf("%d ", root.Val)
}

func preOrderTraversalNonRec(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)

	for root != nil || len(stack) != 0 {
		for root != nil {
			result = append(result, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Right

	}
	return result
}

func inOrderTraversalNonRec(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	stack := make([]*TreeNode, 0)
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		val := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, val.Val)
		root = val.Right

	}
	return result
}

func postOrderTraversalNonRec(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	lastVisit := new(TreeNode)
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		node := stack[len(stack)-1]
		if node.Right == nil || node.Right == lastVisit {
			stack = stack[:len(stack)-1]
			result = append(result, node.Val)
			lastVisit = node
		} else {
			root = node.Right
		}
	}
	return result
}
