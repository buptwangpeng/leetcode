package main

import (
	// "fmt"
	// "leetcode/bfs"
	"fmt"
	"leetcode/window" //文件夹名称必须是英文，不然调用包会报错，找不到
)

func main() {
	// 字符串排列
	s1 := "ac"
	s2 := "abcadedg"
	res := window.CheckInclusion2(s1, s2)
	fmt.Println(res)

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
