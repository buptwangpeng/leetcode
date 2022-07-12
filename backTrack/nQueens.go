package backtrack

var res0 [][]string

// n皇后问题
// 半成品算法，只满足部分测试用例，之后有时间完善

func SolveNQueens(n int) [][]string {
	res0 = make([][]string, 0)
	solution := make([]string, 0)
	backTrack(n, 0, solution)
	// s := ".Q.."
	// fmt.Println("test: ", s[1]=='Q')
	return res0
}

func backTrack(n, row int, solution []string) {
	if row == n {
		// 满足条件就return
		res0 = append(res0, solution)
		solution = make([]string, 0)
		return
	}
	// 初始值
	tmp := make([]byte, 0)
	for j := 0; j < n; j++ {
		tmp = append(tmp, byte('.'))
	}
	// 遍历每一个选择，每一行的每个位置
	for i := 0; i < n; i++ {
		// 做选择
		tmp[i] = byte('Q')
		tmpStr := string(tmp)
		// 如选择无效，则continue
		if !isValid(tmpStr, solution) {
			tmp[i] = byte('.') // 撤销选择，也很很关键
			//    fmt.Printf("tmpStr: %s, solution: %s \n", tmpStr, solution[len(solution)-1])
			continue
		}
		//    if len(solution)!=0 && i==4{
		//        fmt.Printf("tmpStr: %s, solution: %s \n", tmpStr, solution[len(solution)-1])
		//    }
		solution = append(solution, tmpStr)
		backTrack(n, row+1, solution)
		//撤销选择
		tmp[i] = byte('.')
		solution = solution[:len(solution)-1] //左闭右开
	}
}

func isValid(s1 string, solution []string) bool {
	if len(solution) == 0 {
		return true
	}
	// s1为当前的最下面一行
	col := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] == 'Q' {
			col = i
		}
	}

	// 判断同一列是否有皇后
	for j := 0; j < len(solution); j++ {
		if solution[j][col] == 'Q' {
			return false
		}
	}
	// 判断左上对脚线是否有皇后
	j := len(solution) - 1
	h := col - 1
	for j >= 0 && h >= 0 {
		if solution[j][h] == 'Q' {
			return false
		}
		j -= 1
		h -= 1
	}
	// 判断右上
	j = len(solution) - 1
	h = col + 1
	for h < len(s1) && j >= 0 {
		if solution[j][h] == 'Q' {
			return false
		}
		h += 1
		j -= 1
	}

	return true
}
