package main

import (
	"log"
	"reflect"
)

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// 判断两棵二叉搜索树是否一样，即它们各自中序遍历得到的[]int是否一样
func Func(a, b *TreeNode) bool {
	sliceA := traverse(a)
	sliceB := traverse(b)
	log.Printf("sliceA: %+v", sliceA)
	log.Printf("sliceB: %+v", sliceB)
	return reflect.DeepEqual(sliceA, sliceB)
}

// 中序遍历
func traverse(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	left := traverse(root.Left)
	result = append(left, result...)
	result = append(result, root.Value)
	right := traverse(root.Right)
	result = append(result, right...)
	return result
}
