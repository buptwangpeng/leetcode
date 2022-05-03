package doublePointer

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	// 是否有环
	isCycle, interNode := detectIsCycle(head)
	if !isCycle {
		return nil
	}
	// 获取环长度
	fast := interNode
	slow := interNode
	leng := 1
	fast = fast.Next
	for fast != slow {
		fast = fast.Next
		leng += 1
	}
	// 获取入口节点
	slow = head
	fast = head
	for i := 0; i < leng; i++ {
		fast = fast.Next
	}
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

func detectIsCycle(head *ListNode) (bool, *ListNode) {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true, slow
		}
	}
	return false, nil
}
