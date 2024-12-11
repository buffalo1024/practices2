package main

type Node struct {
	Value int
	Next  *Node
}

// 删除链表倒数第n个结点，先计算链表长度，再根据长度和n得出链表正数要删除哪个结点
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

	if backwardsIndexToDelete == length {
		// 倒数第backwardsIndexToDelete个就是链表头结点，删掉头结点
		head = head.Next
		return head
	}

	cur = head
	for i := 1; i < length-backwardsIndexToDelete; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next

	return head
}

// 用两个指针，一个指针先走backwardsIndexToDelete步，然后两个指针一起前进，先走的指针领先后走的指针backwardsIndexToDelete步
// 先走的指针到达链表尾部时，后走的指针也走到了链表倒数backwardsIndexToDelete步附近
func Func2(head *Node, backwardsIndexToDelete int) *Node {
	pointer1 := head
	pointer2 := head
	// pointer1先走backwardsIndexToDelete步
	for i := 1; i <= backwardsIndexToDelete; i++ {
		pointer1 = pointer1.Next
	}

	// pointer1走了backwardsIndexToDelete步后正好到了链表尾部，那么pointer2不用走了，要删除的就是头结点
	if pointer1 == nil {
		head = head.Next
		return head
	}

	for {
		// 如果pointer1已经到达链表尾部，直接删除目标结点并返回头结点
		if pointer1.Next == nil {
			pointer2.Next = pointer2.Next.Next
			return head
		}

		// pointer1未到达链表尾部，两指针继续一同前进
		pointer1 = pointer1.Next
		pointer2 = pointer2.Next
	}
}
