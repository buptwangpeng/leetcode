package others

import "leetcode/utils"

// leetcode 876 简单 链表的中间节点
// 快慢双指针
func MiddleNode(head *utils.ListNode) *utils.ListNode {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}
	return slow
}
