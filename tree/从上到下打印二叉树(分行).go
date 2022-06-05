package tree

import "leetcode/utils"

// leetcode 剑指 Offer 32 - II. 从上到下打印二叉树 II(分行)
// 关键点：队列+两个变量
func LevelOrder2(root *utils.TreeNode) [][]int {
	if root == nil {
		return nil
	}
	// 队列
	queue := make([]*utils.TreeNode, 0)
	resItem := make([]int, 0)
	res := make([][]int, 0)
	queue = append(queue, root)
	// 这道题和「从上到下到下打印二叉树」的区别：
	// 增加两个变量：变量toBePrinted表示在当前层中还没有打印的节点数，变量nextLevel表示下一层的节点数。
	toBePrinted := 1
	nextLevel := 0
	for len(queue) != 0 {
		current := queue[0]
		resItem = append(resItem, current.Value)
		toBePrinted--
		if current.LeftNode != nil {
			queue = append(queue, current.LeftNode)
			nextLevel++
		}
		if current.RightNode != nil {
			queue = append(queue, current.RightNode)
			nextLevel++
		}
		queue = queue[1:]
		if toBePrinted == 0 {
			res = append(res, resItem)
			resItem = make([]int, 0)
			toBePrinted = nextLevel
			nextLevel = 0
		}
	}
	return res
}
