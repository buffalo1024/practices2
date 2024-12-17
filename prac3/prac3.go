package main

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func Func(root *TreeNode) []int {
	result := []int{}
	if root.Left != nil {
		left := Func(root.Left)
		result = append(left, result...)
	}
	result = append(result, root.Value)

	if root.Right == nil {
		return result
	} else {
		right := Func(root.Right)
		result = append(result, right...)
	}

	return result
}
