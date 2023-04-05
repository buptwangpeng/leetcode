package bfs

import (
	"leetcode/utils"
)

func GetTreeMinHeight(root *utils.TreeNode) int {
	// 队列
	q := utils.InitQueue()
	q.Push(root)
	step := 1 // 因为是最小高度，所以初始值为1
	// 树就不需要用visited map，因为不会重复

	// 核心逻辑
	for q.Length() != 0 {
		// 将当前队列的所有节点都向四周扩散
		n := q.Length()
		for i := 0; i < n; i++ {
			current := q.Pop().(*utils.TreeNode)
			// 是否符合目标
			if current.LeftNode == nil && current.RightNode == nil {
				return step
			}
			// 把左右非空子节点放入队列
			if current.LeftNode != nil {
				q.Push(current.LeftNode)
			}
			if current.RightNode != nil {
				q.Push(current.RightNode)
			}
		}
		// 当前队列所有节点都扩散完之后，结果+1
		step += 1
	}

	return step
}

// func main() {
// 	root := utils.InitTree()
// 	res := GetTreeMinHeight(root)
// 	fmt.Println(res)
// }
