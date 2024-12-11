package main

type Node struct {
	Value int
	Next  *Node
}

func Func1(head *Node, backwardsIndexToDelete int) *Node {

	// compute length
	length := 0
	cur := head
	for {
		if cur != nil {
			length = length + 1
			cur = cur.Next
			continue
		} else {
			break
		}
	}

	cur = head
	for i := 1; i < length-backwardsIndexToDelete; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next

	return head
}

func Func2(head *Node, backwardsIndexToDelete int) *Node {
	pointer1 := head
	pointer2 := head
	for i := 1; i <= backwardsIndexToDelete; i++ {
		pointer1 = pointer1.Next
	}
	for {
		pointer1 = pointer1.Next
		pointer2 = pointer2.Next
		if pointer1.Next == nil {
			pointer2.Next = pointer2.Next.Next
			break
		}
	}

	return head
}
