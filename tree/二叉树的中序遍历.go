package tree

import "leetcode/utils"

var res []int

func InorderTraversal(root *utils.TreeNode) []int {
	res = make([]int, 0)
	test(root)
	return res
}
func test(root *utils.TreeNode) {
	if root == nil {
		return
	}
	test(root.LeftNode)
	res = append(res, root.Value)
	test(root.RightNode)
}
