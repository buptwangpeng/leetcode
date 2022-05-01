package main

import (
	"fmt"
	"leetcode/bfs"
)

func main() {
	// 滑动谜题
	board := [][]int{{4, 1, 2}, {5, 0, 3}}
	res := bfs.SlidingPuzzle(board)
	fmt.Println(res) //5

	// 二叉树的最小高度
	// root := utils.InitTree()
	// res := bfs.GetTreeMinHeight(root)
	// fmt.Println(res) //2

	// 打开转盘锁的最小次数
	// deadArr := []string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}
	// res := bfs.GetOpenLockMin("8888", deadArr)
	// fmt.Println(res) //-1
}
