package tree

import "leetcode/utils"

// leetcode 剑指 Offer 32 - III. 从上到下打印二叉树 III

//关键：两个栈（写进一个数组里)，特殊的写法，最后还反了一下

func LevelOrder3(root *utils.TreeNode) [][]int {
	if root == nil {
		return nil
	}
	// 两个栈，不用队列。后进先出
	stack := make([][]*utils.TreeNode, 0)
	for i := 0; i < 2; i++ {
		temp := make([]*utils.TreeNode, 0)
		stack = append(stack, temp)
	}
	resItem := make([]int, 0)
	res := make([][]int, 0)
	//
	current := 0
	next := 1
	//
	stack[current] = append(stack[current], root)
	for len(stack[current]) != 0 {
		currentNode := stack[current][len(stack[current])-1]
		resItem = append(resItem, currentNode.Value)
		stack[current] = stack[current][:len(stack[current])-1] // 左闭右开, 栈pop
		if current == 0 {
			if currentNode.LeftNode != nil {
				stack[next] = append(stack[next], currentNode.LeftNode)
			}
			if currentNode.RightNode != nil {
				stack[next] = append(stack[next], currentNode.RightNode)
			}
		} else {
			// 关键：此处还反了一下注意呀
			if currentNode.RightNode != nil {
				stack[next] = append(stack[next], currentNode.RightNode)
			}
			if currentNode.LeftNode != nil {
				stack[next] = append(stack[next], currentNode.LeftNode)
			}
		}

		if len(stack[current]) == 0 {
			current, next = next, current
			res = append(res, resItem)
			resItem = make([]int, 0)
		}

	}
	return res

}
