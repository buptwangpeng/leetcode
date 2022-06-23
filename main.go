package main

import (
	"fmt"
	"leetcode/tree"
	"leetcode/utils"
)

//文件夹名称必须是英文，不然调用包会报错，找不到

func main() {
	// 完全二叉树的节点数
	root := utils.InitTree()
	res := tree.CountNums(root)
	fmt.Println(res)
	// 0-1背包
	// res := dynamic.Knapsack(4, []int{2, 1, 3}, []int{4, 2, 3})
	// fmt.Println(res)
	// 从上到下打印二叉树(不分行)

	// 对称的二叉树

	// 二叉树的镜像

	// 树的子结构

	// 二叉搜索树中的中序后继

	// 恢复二叉搜索树

	// 从前序与中序遍历序列构造二叉树
	// pre := []int{3, 9, 20, 15, 7}
	// mid := []int{9, 3, 15, 20, 7}
	// root := tree.BuildTree(pre, mid)
	// fmt.Println(root.Value)

	// 二叉树中的最大路径和
	// root := utils.InitTree()
	// res := tree.MaxPathSum(root)
	// fmt.Println(res) //8

	// 搜索二维矩阵
	// test := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	// res := binarysearch.SearchMatrix(test, 10)
	// fmt.Println("res: ", res)
	// x的平方根
	// res := binarysearch.MySqrt(8)
	// fmt.Println("res: ", res)
	// 在排序数组中查找元素的第一个和最后一个位置
	// test := []int{1, 2, 4, 4, 4, 5, 6}
	// res := binarysearch.SearchRange(test, 4)
	// fmt.Println("res: ", res)

	// 搜索旋转排序数组
	// test := []int{4, 5, 6, 0, 1, 2, 3}
	// res := binarysearch.Search(test, 5)
	// fmt.Println("res: ", res)
	// n皇后
	// res := backtrack.SolveNQueens(5) //6还有点问题，半成品算法
	// fmt.Println(res)

	// 无重复全排列
	// nums := []int{1, 2, 3}
	// res := backtrack.Permute(nums)
	// fmt.Println(res)

	// 最长无重复子串
	// s := "cbaebabacd"
	// res := window.LengthOfLongestSubstring(s)
	// fmt.Println(res)
	// 找所有字母异位词
	// s1 := "cbaebabacd"
	// s2 := "abc"
	// res := window.FindAnagrams(s1, s2)
	// fmt.Println(res)

	// 字符串排列
	// s1 := "ac"
	// s2 := "abcadedg"
	// res := window.CheckInclusion2(s1, s2)
	// fmt.Println(res)

	//最小覆盖子串
	// s := "ADOBECODEBANC"
	// target := "ABC"
	// res0 := window.MinWindow(s, target)
	// fmt.Println(res0)

	// 滑动谜题
	// board := [][]int{{4, 1, 2}, {5, 0, 3}}
	// res := bfs.SlidingPuzzle(board)
	// fmt.Println(res) //5

	// 二叉树的最小高度
	// root := utils.InitTree()
	// res := bfs.GetTreeMinHeight(root)
	// tree.InorderTraversal(root)
	// fmt.Println(res) //2

	// 打开转盘锁的最小次数
	// deadArr := []string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}
	// res := bfs.GetOpenLockMin("8888", deadArr)
	// fmt.Println(res) //-1
}
