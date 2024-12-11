package main

import (
	"testing"
)

var (
	// test1IntSlicesAndIndexToDelete = IntSlicesAndIndexToDelete{
	// 	Slice:                  []int{1, 2, 3, 4, 5, 6, 7},
	// 	BackwardsIndexToDelete: 3,
	// }
	testcase1 = TestCase{
		Slice:                  []int{1, 2, 3, 4, 5, 6, 7},
		BackwardsIndexToDelete: 3,
		WantedResultSlice:      []int{1, 2, 3, 4, 6, 7},
	}

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
	}
)

type IntSlicesAndIndexToDelete struct {
	Slice                  []int
	BackwardsIndexToDelete int
}

type TestCase struct {
	Slice                  []int
	Head                   *Node
	BackwardsIndexToDelete int
	WantedResultSlice      []int
}

func TestFunc1(t *testing.T) {
	for k, tc := range testCases {
		t.Logf("starting testcase %v", k)
		testCase := generateTestCase(tc.Slice, tc.WantedResultSlice, tc.BackwardsIndexToDelete)
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

	// testCase := generateTestCase(testcase1.Slice, testcase1.WantedResultSlice, testcase1.BackwardsIndexToDelete)
	// got := Func1(testCase.Head, testCase.BackwardsIndexToDelete)
	// cur := got
	// i := 0
	// for {
	// 	if cur != nil {
	// 		t.Logf("value: %v, want: %v", cur.Value, testCase.WantedResultSlice[i])
	// 		if cur.Value != testCase.WantedResultSlice[i] {
	// 			t.Fail()
	// 		}
	// 		cur = cur.Next
	// 		i = i + 1
	// 	} else {
	// 		break
	// 	}
	// }
}

func TestFunc2(t *testing.T) {
	testCase := generateTestCase(testcase1.Slice, testcase1.WantedResultSlice, testcase1.BackwardsIndexToDelete)
	got := Func2(testCase.Head, testCase.BackwardsIndexToDelete)
	cur := got
	i := 0
	for {
		if cur != nil {
			t.Logf("value: %v, want: %v", cur.Value, testCase.WantedResultSlice[i])
			if cur.Value != testCase.WantedResultSlice[i] {
				t.Fail()
			}
			cur = cur.Next
			i = i + 1
		} else {
			break
		}
	}
}

func generateTestCase(intSlice, wantedResultSlice []int, indexToDelete int) TestCase {
	return TestCase{
		Head:                   transformIntSliceToChain(intSlice),
		BackwardsIndexToDelete: indexToDelete,
		WantedResultSlice:      wantedResultSlice,
	}
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
