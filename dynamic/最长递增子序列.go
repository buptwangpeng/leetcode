package dynamic

// leetcode 300 最长递增子序列

// 动态规划
// 步骤：找到“状态”和“选择”--》明确dp数组/函数的定义--〉寻找状态之间的关系
// dp[i]表示以num[i]这个数结尾的最长递增子序列的长度
func LengthOfLIS(nums []int) int {
	// 初始化
	dp := make([]int, 0)
	for n := 0; n < len(nums); n++ {
		dp = append(dp, 1)
	}
	// dp[i]状态转化
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	//取dp数组中的最大值，遍历
	res := 0
	for _, v := range dp {
		if v > res {
			res = v
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
