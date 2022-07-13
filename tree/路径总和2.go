package tree

import (
	"fmt"
	"leetcode/utils"
)

// leetcode 113. 路径总和 II 中等
// 通过105/115
var sols [][]int

func pathSum(root *utils.TreeNode, targetSum int) [][]int {
	sols = make([][]int, 0)
	sol := make([]int, 0)
	test0(root, targetSum, sol)
	fmt.Println("res_sols: ", sols)
	return sols
}

func test0(root *utils.TreeNode, targetSum int, sol []int) {
	if root == nil {
		return
	}
	sol = append(sol, root.Value)
	if root.LeftNode == nil && root.RightNode == nil && root.Value == targetSum {
		sols = append(sols, sol)
		fmt.Println("sols: ", sols)
		return
	}
	test0(root.LeftNode, targetSum-root.Value, sol)
	test0(root.RightNode, targetSum-root.Value, sol)
	return
}
