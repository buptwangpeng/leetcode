package main

import (
	"fmt"
	backtrack "leetcode/backTrack"
)

//文件夹名称必须是英文，不然调用包会报错，找不到

func main() {
	// n皇后
	res := backtrack.SolveNQueens(5) //6还有点问题，半成品算法
	fmt.Println(res)

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
