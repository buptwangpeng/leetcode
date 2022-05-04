package window

// s2中是否有一个子串：只包含s2的全部字符，而没有其他字符。用滑动窗口

// 方法1 我自己写的，leetcode上可以过
func CheckInclusion1(s1 string, s2 string) bool {
	// 初始化：两个map，三个int变量
	need := make(map[uint8]int)
	window := make(map[uint8]int)
	left := 0
	right := 0
	valid := 0 //很重要
	// need 初始化
	for i := 0; i < len(s1); i++ {
		need[s1[i]] += 1
	}
	// 子串的长度
	leng := len(s2) + 1000 // 长度比s2大即可
	// 循环
	for right < len(s2) {
		// c为将为写入窗口的右侧元素
		c := s2[right]
		// 窗口右移
		right += 1
		// 数据更新的逻辑
		_, ok := need[c]
		if ok {
			window[c] += 1
			if window[c] == need[c] {
				valid += 1
			}
		}
		// debug部分
		// fmt.Printf("window: [%d, %d) \n", left, right)

		// 判断窗口是否要左侧收缩
		for valid == len(need) {
			// 更新子串的长度
			if right-left < leng {
				leng = right - left
			}
			// d为将要移出窗口的左侧元素
			d := s2[left]
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
	if leng == len(s1) {
		return true
	}
	return false
}

// 方法2：labuladong书上的
// s2中是否有一个子串：只包含s2的全部字符，而没有其他字符。用滑动窗口
func CheckInclusion2(s1 string, s2 string) bool {
	// 初始化：两个map，三个int变量
	need := make(map[uint8]int)
	window := make(map[uint8]int)
	left := 0
	right := 0
	valid := 0 //很重要
	// need 初始化
	for i := 0; i < len(s1); i++ {
		need[s1[i]] += 1
	}
	// 循环
	for right < len(s2) {
		// c为将为写入窗口的右侧元素
		c := s2[right]
		// 窗口右移
		right += 1
		// 数据更新的逻辑
		_, ok := need[c]
		if ok {
			window[c] += 1
			if window[c] == need[c] {
				valid += 1
			}
		}
		// debug部分
		// fmt.Printf("window: [%d, %d) \n", left, right)

		// 判断窗口是否要左侧收缩
		for right-left >= len(s1) {
			if valid == len(need) {
				return true
			}
			// d为将要移出窗口的左侧元素
			d := s2[left]
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
	return false
}
