package main

import (
	"testing"
)

var (
	node1  = &TreeNode{Value: 1}
	node2  = &TreeNode{Value: 2}
	node3  = &TreeNode{Value: 3}
	node4  = &TreeNode{Value: 4}
	node5  = &TreeNode{Value: 5}
	node6  = &TreeNode{Value: 6}
	node7  = &TreeNode{Value: 7}
	node8  = &TreeNode{Value: 8}
	node9  = &TreeNode{Value: 9}
	wanted = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
)

func makeTreeA() *TreeNode {
	n1 := &TreeNode{Value: node1.Value}
	n2 := &TreeNode{Value: node2.Value}
	n3 := &TreeNode{Value: node3.Value}
	n4 := &TreeNode{Value: node4.Value}
	n5 := &TreeNode{Value: node5.Value}
	n6 := &TreeNode{Value: node6.Value}
	n7 := &TreeNode{Value: node7.Value}
	n8 := &TreeNode{Value: node8.Value}
	n9 := &TreeNode{Value: node9.Value}

	n6.Left = n2
	n2.Left = n1
	n2.Right = n3
	n6.Right = n7
	n7.Right = n9
	n9.Left = n8
	n3.Right = n4
	n4.Right = n5
	return n6
}

func makeTreeB() *TreeNode {
	n1 := &TreeNode{Value: node1.Value}
	n2 := &TreeNode{Value: node2.Value}
	n3 := &TreeNode{Value: node3.Value}
	n4 := &TreeNode{Value: node4.Value}
	n5 := &TreeNode{Value: node5.Value}
	n6 := &TreeNode{Value: node6.Value}
	n7 := &TreeNode{Value: node7.Value}
	n8 := &TreeNode{Value: node8.Value}
	n9 := &TreeNode{Value: node9.Value}

	n6.Left = n2
	n2.Left = n1
	n2.Right = n3
	n6.Right = n7
	n7.Right = n9
	n9.Left = n8
	n3.Right = n4
	n7.Left = n5
	return n6
}

func TestFunc(t *testing.T) {
	treeA1 := makeTreeA()
	treeA2 := makeTreeA()
	treeB := makeTreeB()
	if !Func(treeA1, treeA2) {
		t.Logf("treeA1 not the same as treeA2")
		t.Fail()
	}
	if Func(treeA1, treeB) {
		t.Logf("treeA1 and treeB are same")
		t.Fail()
	}
}
