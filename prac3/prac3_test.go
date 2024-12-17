package main

import (
	"reflect"
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

func makeTree() *TreeNode {
	node6.Left = node2
	node2.Left = node1
	node2.Right = node3
	node6.Right = node7
	node7.Right = node9
	node9.Left = node8
	node3.Right = node4
	node4.Right = node5
	return node6
}

func TestFunc(t *testing.T) {
	makeTree()
	got := Func(node6)
	t.Logf("got: %+v", got)
	if !reflect.DeepEqual(got, wanted) {
		t.Fail()
	}

}
