package binarysearch

// leetcode 34 中等
func SearchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	leftBoard := leftBoard(nums, target)
	rightBoard := rightBoard(nums, target)
	return []int{leftBoard, rightBoard}
}

func leftBoard(nums []int, target int) int {
	// 默写模版
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if target < nums[mid] {
			right = mid - 1
		} else if target > nums[mid] {
			left = mid + 1
		} else if target == nums[mid] {
			// 从右侧向左缩进
			right = mid - 1
		}
	}
	if left >= len(nums) || nums[left] != target {
		return -1
	}

	return left
}
func rightBoard(nums []int, target int) int {
	// 默写模版
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if target < nums[mid] {
			right = mid - 1
		} else if target > nums[mid] {
			left = mid + 1
		} else if target == nums[mid] {
			// 从左侧向右缩进
			left = mid + 1
		}
	}
	if right <= -1 || nums[right] != target {
		return -1
	}

	return right
}
