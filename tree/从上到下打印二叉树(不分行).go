package tree

import "leetcode/utils"

// leetcode 剑指 Offer 32 - I. 从上到下打印二叉树(不分行)

func LevelOrder(root *utils.TreeNode) []int {
	if root == nil {
		return nil
	}
	// 要用到队列
	queue := make([]*utils.TreeNode, 0) // 数组模仿队列
	res := make([]int, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		current := queue[0]
		// 核心：取出队列头部的节点打印，若该节点有子节点就把子节点写入队列
		res = append(res, current.Value)
		if current.LeftNode != nil {
			queue = append(queue, current.LeftNode)
		}
		if current.RightNode != nil {
			queue = append(queue, current.RightNode)
		}
		// 先进先出，队列
		queue = queue[1:]
	}
	return res
}
