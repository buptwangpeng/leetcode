package window

func MinWindow(s string, t string) string {
	// 初始化：两个map，三个int变量
	window := make(map[uint8]int)
	need := make(map[uint8]int)
	left := 0
	right := 0
	valid := 0 // 有效的子字符数量，无重复
	for i := 0; i < len(t); i++ {
		need[t[i]] += 1
	}
	// 最小覆盖子串的起始索引及长度
	start := 0
	leng := len(s) + 1 // 长度初始值值设的比s的长度大一点即可
	//
	for right < len(s) {
		// c为将移入窗口的字符，go中c的类型是uint8
		c := s[right]
		// 右移窗口
		right += 1
		// 进行窗口内数据的一系列更新
		_, ok := need[c]
		if ok {
			window[c] += 1
			if window[c] == need[c] {
				valid += 1
			}
		}

		// debug
		// fmt.Printf("window: [%d, %d) \n", left, right)
		// fmt.Println("valid: ", valid)
		// fmt.Println("need: ", need)

		// 判断左侧窗口是否要收缩
		for valid == len(need) { //这里很关键，必须用need的长度，不然测试用例: s="aa", t="aa"就会报错
			//  更新最小覆盖子串
			if right-left < leng {
				start = left
				leng = right - left
			}
			// d为将移出窗口的字符
			d := s[left]
			// 左移窗口
			left += 1
			// 进行窗口内的一系列更新
			_, ok := need[d]
			if ok {
				if window[d] == need[d] {
					valid -= 1
				}
				window[d] -= 1
			}
		}
	}

	if leng == len(s)+1 {
		return ""
	}
	return s[start : start+leng] //左闭右开

}
