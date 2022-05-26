package tree

import "leetcode/utils"

//leetcode 124 困难

var ans int

func MaxPathSum(root *utils.TreeNode) int {
	ans = root.Value
	maxPathSum0(root)
	return ans
}

func maxPathSum0(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}
	left := max(0, maxPathSum0(root.LeftNode)) //这样写的原因：题目是从树的任意结点出发
	right := max(0, maxPathSum0(root.RightNode))
	/***后序遍历***/
	ans = max(ans, left+right+root.Value)
	return max(left, right) + root.Value
	/***/
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
