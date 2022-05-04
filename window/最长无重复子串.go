package window

func LengthOfLongestSubstring(s string) int {
	// 初始化
	window := make(map[uint8]int)
	left := 0
	right := 0
	// 返回
	res := 0
	// 循环
	for right < len(s) {
		// c为将从右侧写入窗口的元素
		c := s[right]
		// 窗口右移
		right += 1
		// 数据更新
		window[c] += 1
		// 判断窗口是否要左侧收缩
		for window[c] > 1 {
			//d为从左侧从窗口中去除的元素
			d := s[left]
			// 窗口左缩
			left += 1
			// 数据更新
			window[d] -= 1
		}
		// 此时[left: right+1]之间的元素是不含重复字符的子串
		res = maxTest(res, right-left)
	}
	return res
}

func maxTest(a, b int) int {
	if a > b {
		return a
	}
	return b
}
