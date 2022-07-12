package binarysearch

// leetcode 74 搜索二维矩阵

func SearchMatrix(matrix [][]int, target int) bool {
	// 把二维数组变为一个递增数组
	nums := make([]int, 0)
	for i := 0; i < len(matrix); i++ {
		nums = append(nums, matrix[i]...)
	}
	// 一般的二分搜索
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if target < nums[mid] {
			right = mid - 1
		} else if target > nums[mid] {
			left = mid + 1
		} else if target == nums[mid] {
			return true
		}
	}
	return false
}
