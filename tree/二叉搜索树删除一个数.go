package tree

import "leetcode/utils"

// leetcode 中等 450 二叉搜索树删除一个数

func deleteNode(root *utils.TreeNode, key int) *utils.TreeNode {
	if root == nil {
		return nil
	}
	// 背模版
	if root.Value == key {
		// 这两个if把情况1和2都正确处理了
		if root.LeftNode == nil {
			return root.RightNode
		} else if root.RightNode == nil {
			return root.LeftNode
		}
		// 处理情况3:找右子树的最小节点(找左子树的最大节点也可)
		minNode := getMinNode(root.RightNode)
		// 把root改成minNode
		root.Value = minNode.Value
		// 转而去删除minNode
		root.RightNode = deleteNode(root.RightNode, minNode.Value)

	} else if root.Value > key {
		root.LeftNode = deleteNode(root.LeftNode, key)
	} else if root.Value < key {
		root.RightNode = deleteNode(root.RightNode, key)
	}
	return root
}
func getMinNode(root *utils.TreeNode) *utils.TreeNode {
	for root.LeftNode != nil {
		root = root.LeftNode
	}
	return root
}
