package tree

import "leetcode/utils"

// leetcode 剑指 Offer 28. 对称的二叉树

// 方法2
func IsSymmetric(root *utils.TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetric0(root.LeftNode, root.RightNode)
}

func isSymmetric0(left, right *utils.TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Value != right.Value {
		return false
	}
	// 该层符合对称二叉树的要求，开始比较下一层
	return isSymmetric0(left.LeftNode, right.RightNode) && isSymmetric0(left.RightNode, right.LeftNode)
}

// 方法1: 比较原树和镜像树的前序遍历 通过测试用例：72 / 195
// func isSymmetric(root*TreeNode) bool {

//     mTree := mirrorTree(root)
//     return isSymmetric0(root, mTree)
// }

// func isSymmetric0(root, rootMirror *TreeNode) bool {
//     if root==nil && rootMirror==nil{
//         return true
//     }
//     if root==nil || rootMirror==nil{
//         return false
//     }
//     if root.Val != rootMirror.Val{
//         return false
//     }
//     left:=isSymmetric0(root.Left, rootMirror.Left)
//     right:=isSymmetric0(root.Right, rootMirror.Right)
//     return left && right
// }

// func mirrorTree(root *TreeNode)*TreeNode{
//     if root==nil{
//         return nil
//     }
//     root.Left, root.Right = root.Right, root.Left
//     mirrorTree(root.Left)
//     mirrorTree(root.Right)
//     return root
// }
