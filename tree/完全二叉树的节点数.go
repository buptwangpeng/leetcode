package tree

import (
	"leetcode/utils"
	"math"
)

// 完全二叉树和满二叉树的定义（中文定义，可看labuladong）
func CountNums(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}
	var hLeft, hRight int
	nLeft := root
	nRight := root
	for nLeft != nil {
		nLeft = nLeft.LeftNode
		hLeft += 1
	}
	for nRight != nil {
		nRight = nRight.RightNode
		hRight += 1
	}
	// 判断是不是满二叉树
	if hLeft == hRight {
		return int(math.Pow(2, float64(hLeft))) - 1
	}
	// 是一般的树
	return 1 + CountNums(root.LeftNode) + CountNums(root.RightNode)
}
