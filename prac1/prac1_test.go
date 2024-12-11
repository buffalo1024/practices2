package main

import (
	"testing"
)

var (
	testCases = []TestCase{
		{
			Slice:                  []int{1, 2, 3, 4, 5, 6, 7},
			BackwardsIndexToDelete: 3,
			WantedResultSlice:      []int{1, 2, 3, 4, 6, 7},
		},
		{
			Slice:                  []int{1, 2, 3, 4, 5, 6, 7},
			BackwardsIndexToDelete: 1,
			WantedResultSlice:      []int{1, 2, 3, 4, 5, 6},
		},
		{
			Slice:                  []int{1, 2, 3, 4, 5, 6, 7},
			BackwardsIndexToDelete: 7,
			WantedResultSlice:      []int{2, 3, 4, 5, 6, 7},
		},
		{
			Slice:                  []int{1, 2},
			BackwardsIndexToDelete: 1,
			WantedResultSlice:      []int{1},
		},
		{
			Slice:                  []int{1, 2},
			BackwardsIndexToDelete: 2,
			WantedResultSlice:      []int{2},
		},
		{
			Slice:                  []int{1},
			BackwardsIndexToDelete: 1,
			WantedResultSlice:      []int{},
		},
	}
)

type TestCase struct {
	Slice                  []int //用于构建链表的序列
	Head                   *Node //链表的头结点
	BackwardsIndexToDelete int   //删除链表倒数第n个结点，此为n
	WantedResultSlice      []int //处理后的链表序列化为slice以检验
}

func TestFunc1(t *testing.T) {
	for k, testCase := range testCases {
		t.Logf("starting testcase %v", k)
		testCase.Head = transformIntSliceToChain(testCase.Slice)
		got := Func1(testCase.Head, testCase.BackwardsIndexToDelete)
		cur := got
		i := 0
	check:
		for {
			if cur != nil {
				t.Logf("value: %v, want: %v", cur.Value, testCase.WantedResultSlice[i])
				if cur.Value != testCase.WantedResultSlice[i] {
					t.Fail()
				}
				cur = cur.Next
				i = i + 1
			} else {
				break check
			}
		}
	}
}

func TestFunc2(t *testing.T) {
	for k, testCase := range testCases {
		t.Logf("starting testcase %v", k)
		testCase.Head = transformIntSliceToChain(testCase.Slice)
		got := Func2(testCase.Head, testCase.BackwardsIndexToDelete)
		cur := got
		i := 0
	check:
		for {
			if cur != nil {
				t.Logf("value: %v, want: %v", cur.Value, testCase.WantedResultSlice[i])
				if cur.Value != testCase.WantedResultSlice[i] {
					t.Fail()
				}
				cur = cur.Next
				i = i + 1
			} else {
				break check
			}
		}
	}
}

func fulfillTestCase(c TestCase) TestCase {
	c.Head = transformIntSliceToChain(c.Slice)
	return c
}

func transformIntSliceToChain(intSlice []int) (head *Node) {
	if len(intSlice) == 0 {
		return
	}

	var former *Node
	for k, v := range intSlice {
		node := &Node{
			Value: v,
		}
		if k == 0 {
			head = node
		}
		if former != nil {
			former.Next = node
		}
		former = node
	}
	return
}
