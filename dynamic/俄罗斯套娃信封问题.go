package dynamic

// leetcode 354 俄罗斯套娃信封问题

// 信封嵌套问题实际上是最长递增子序列问题上升到二维
// 关键：先对宽度w进行升序排序，如果遇到w相同的情况,则按照高度h降序排序（为了让w相同的信封只选一个）。之后把所有的h作为一个数组，在这个数组上计算出的LIS的长度就是答案
import "sort"

func MaxEnvelopes(envelopes [][]int) int {
	// 排序
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})
	// 对h进行最长递增子序列
	data := make([]int, 0)
	for i := 0; i < len(envelopes); i++ {
		data = append(data, envelopes[i][1])
	}
	return LengthOfLIS(data)
}
