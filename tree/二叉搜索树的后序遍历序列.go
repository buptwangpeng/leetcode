package tree

// leetcode 剑指 Offer 33. 二叉搜索树的后序遍历序列
func VerifyPostorder(postorder []int) bool {
	if postorder == nil || len(postorder) == 0 {
		return true
	}
	root := postorder[len(postorder)-1]
	tag := false
	res := true

	leftArr := make([]int, 0)
	rightArr := make([]int, 0)
	for i := 0; i < len(postorder)-1; i++ {
		if postorder[i] < root && tag {
			res = false
		}
		if postorder[i] < root && !tag {
			leftArr = append(leftArr, postorder[i])
		}
		if postorder[i] > root {
			rightArr = append(rightArr, postorder[i])
			tag = true
		}

	}

	left := VerifyPostorder(leftArr)
	right := VerifyPostorder(rightArr)
	return left && right && res
}
