package tree

import (
	"leetcode/utils"
	"strconv"
)

// leetcode 简单 257. 二叉树的所有路径
// 通过测试用例： 187 / 208 (经我测试，感觉力扣有bug，要不就是我的全局变量使用有问题，类似的问题出现过)

var res11 [][]string

func binaryTreePaths(root *utils.TreeNode) []string {
	res11 = make([][]string, 0)
	s := make([]string, 0)
	binaryTreePaths0(root, s)
	//  fmt.Println("res: ", res)
	// 格式化返回
	return formatRes(res11)
}
func formatRes(strArr [][]string) []string {
	res := make([]string, 0)
	for _, v := range strArr {
		s := ""
		for i := 0; i < len(v); i++ {
			s += v[i]
			if i != len(v)-1 {
				s += "->"
			}
		}
		res = append(res, s)
	}
	return res
}

func binaryTreePaths0(root *utils.TreeNode, s []string) {
	// 先背模版
	if root == nil {
		return
	}
	valStr := strconv.Itoa(root.Value)
	s = append(s, valStr)
	if root.LeftNode == nil && root.RightNode == nil {
		//  fmt.Println(" ")
		//  fmt.Println("item: ", s)
		res11 = append(res11, s)
	}
	binaryTreePaths0(root.LeftNode, s)
	binaryTreePaths0(root.RightNode, s)
}
