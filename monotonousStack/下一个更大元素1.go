package monotonousstack

//leetcode 496 easy // 下一个更大元素1

// 单调栈+map
// 先求 nums2中全部元素的下一个更大元素，再从中取nums1对应的值
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	// nums2 的单调栈
	res := make(map[int]int)
	stack := make([]int, 0)
	// 倒着写入栈
	for i := len(nums2) - 1; i >= 0; i-- {
		//
		for len(stack) != 0 && nums2[i] >= stack[len(stack)-1] {
			// 出栈，小个子
			stack = stack[:len(stack)-1]
		}

		//结果赋值
		if len(stack) != 0 {
			res[nums2[i]] = stack[len(stack)-1]
		} else {
			res[nums2[i]] = -1
		}
		// 入栈：进行身高比较
		stack = append(stack, nums2[i])
	}
	// 返回结果
	res0 := make([]int, 0)
	for i := 0; i < len(nums1); i++ {
		res0 = append(res0, res[nums1[i]])
	}
	return res0
}
