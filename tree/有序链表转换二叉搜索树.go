package tree

import "leetcode/utils"

// leetcode  109 有序链表转换二叉搜索树 中等

func sortedListToBST(head *utils.ListNode) *utils.TreeNode {
	return buildTree(head, nil) // 右边界为nil
}

func getMid(left, right *utils.ListNode) *utils.ListNode {
	fast, slow := left, left
	for fast != right && fast.Next != right {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func buildTree(left, right *utils.ListNode) *utils.TreeNode {
	// 这个条件很关键，原因是右边界是nil
	if left == right {
		return nil
	}
	mid := getMid(left, right)
	root := &utils.TreeNode{mid.Val, nil, nil}
	root.LeftNode = buildTree(left, mid)
	root.RightNode = buildTree(mid.Next, right)
	return root
}
