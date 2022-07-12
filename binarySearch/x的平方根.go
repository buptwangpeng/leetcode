package binarysearch

// leetcode 79 简单 x的平方根
func MySqrt(x int) int {
	left := 0
	right := x
	ans := -1
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid <= x {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}
