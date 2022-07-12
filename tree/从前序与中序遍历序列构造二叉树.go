package tree

import "leetcode/utils"

// leetcode 105 中等 从前序与中序遍历序列构造二叉树

func BuildTree(preorder []int, inorder []int) *utils.TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &utils.TreeNode{preorder[0], nil, nil}
	index := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			index = i
			break
		}
	}

	root.LeftNode = BuildTree(preorder[1:len(inorder[:index])+1], inorder[:index])
	root.RightNode = BuildTree(preorder[len(inorder[:index])+1:], inorder[index+1:])
	return root
}
