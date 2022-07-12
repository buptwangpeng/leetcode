package tree

import "leetcode/utils"

// leetcode 108 简单 将有序数组转换为二叉搜索树

// 二叉搜索树的中序遍历是升序的；本题要求高度平衡，因此我们需要选择升序序列的中间元素作为根节点
func SortedArrayToBST(nums []int) *utils.TreeNode {
	return bst(nums, 0, len(nums)-1)
}
func bst(nums []int, left, right int) *utils.TreeNode {
	if right < left {
		return nil
	}
	mid := (left + right) / 2
	root := &utils.TreeNode{Value: nums[mid]}
	root.LeftNode = bst(nums, left, mid-1)
	root.RightNode = bst(nums, mid+1, right)
	return root
}
