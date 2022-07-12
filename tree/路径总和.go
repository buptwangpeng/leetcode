package tree

import "leetcode/utils"

// leetcode 简单 112 路径总和

//方法1:
func hasPathSum(root *utils.TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.LeftNode == nil && root.RightNode == nil {
		return targetSum == root.Value
	}
	return hasPathSum(root.LeftNode, targetSum-root.Value) || hasPathSum(root.RightNode, targetSum-root.Value)
}

// 方法2:通过率105/117
// func hasPathSum(root *TreeNode, targetSum int) bool {
//     if root==nil{
//         return false
//     }
//     return test(root,targetSum,0)
// }
// func test(root *TreeNode, targetSum, currentSum int)bool{
//     if root==nil{
//         if currentSum== targetSum{
//             return true
//         }
//         return false
//     }
//     currentSum = currentSum+root.Val

//     left := test(root.Left,targetSum,currentSum)
//     right := test(root.Right, targetSum, currentSum)
//     return left || right
// }
