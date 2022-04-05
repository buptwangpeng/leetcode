package utils

// mock 测试二叉树
type TreeNode struct {
	Value     int
	LeftNode  *TreeNode
	RightNode *TreeNode
}

func InitTree() *TreeNode {
	// a5 := &TreeNode{
	// 	Value:     5,
	// 	LeftNode:  nil,
	// 	RightNode: nil,
	// }
	a4 := &TreeNode{
		Value:     4,
		LeftNode:  nil,
		RightNode: nil,
	}
	a3 := &TreeNode{
		Value:     3,
		LeftNode:  nil,
		RightNode: nil,
	}
	a2 := &TreeNode{
		Value:     2,
		LeftNode:  nil,
		RightNode: nil,
	}
	a1 := &TreeNode{
		Value:     1,
		LeftNode:  a3,
		RightNode: a4,
	}
	a0 := &TreeNode{
		Value:     0,
		LeftNode:  a1,
		RightNode: a2,
	}
	return a0
}
