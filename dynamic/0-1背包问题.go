package dynamic

// 0-1背包问题
/***
模版：
for 状态1 in 状态1所有的取值：
   for 状态2 in 状态1所有的取值：
	  for ...
	    dp[状态1][状态2][...] = 择优(选择1，选择2，...)
***/
func Knapsack(w int, weights, values []int) int {
	if len(values) != len(weights) {
		return -1
	}
	// 初始化
	n := len(values)
	dp := make([][]int, 0)
	for i := 0; i <= n; i++ {
		dp0 := make([]int, 0)
		for j := 0; j <= w; j++ {
			dp0 = append(dp0, 0)
		}
		dp = append(dp, dp0)
	}

	// 状态转移
	for i := 1; i <= n; i++ {
		for j := 1; j <= w; j++ {
			if j < weights[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j-weights[i-1]]+values[i-1], dp[i-1][j])
			}
		}
	}
	return dp[n][w]
}
