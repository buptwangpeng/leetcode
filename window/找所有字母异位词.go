package window

// leetcode 438
// 异位词其实就是全排列

func FindAnagrams(s string, p string) []int {
	// 初始化
	need := make(map[uint8]int)
	window := make(map[uint8]int)
	left := 0
	right := 0
	valid := 0
	// need map 初始化
	for i := 0; i < len(p); i++ {
		need[p[i]] += 1
	}
	//返回
	res := make([]int, 0)
	//循环
	for right < len(s) {
		// c为将要从右侧加入窗口的元素
		c := s[right]
		// 窗口右移
		right += 1
		// 数据更新
		_, ok := need[c]
		if ok {
			window[c] += 1
			if window[c] == need[c] {
				valid += 1
			}
		}
		// debug
		// fmt.Printf("window [%d, %d) \n", left, right)

		// 判断是否要窗口左缩
		for right-left >= len(p) {
			// 满足条件返回
			if valid == len(need) {
				res = append(res, left)
			}
			// d为将要从左侧移出窗口的元素
			d := s[left]
			// 窗口左缩
			left += 1
			// 数据更新
			_, ok := need[d]
			if ok {
				if window[d] == need[d] {
					valid -= 1
				}
				window[d] -= 1
			}
		}
	}

	return res
}
