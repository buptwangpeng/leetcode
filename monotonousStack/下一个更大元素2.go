package monotonousstack

// leetcode 中等 下一个最大元素-循环数组

func nextGreaterElements2(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	stack := make([]int, 0)
	// 假装nums长度为2倍
	for i := 2*n - 1; i >= 0; i-- {
		// 比身高循环
		for len(stack) != 0 && nums[i%n] >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		//处理结果
		if len(stack) != 0 {
			res[i%n] = stack[len(stack)-1]
		} else {
			res[i%n] = -1
		}
		//入栈
		stack = append(stack, nums[i%n])
	}
	return res
}
