package main

import (
	"fmt"
	"leetcode/utils"
)

/*
打开转盘锁， 《labuladong的算法小抄》1.4.3
*/
func GetOpenLockMin(target string, deadArr []string) int {
	// 队列
	q := utils.InitQueue()
	// 记录走过的路径，避免重复
	visited := make(map[string]bool)
	deadMap := make(map[string]bool)
	for i := 0; i < len(deadArr); i++ {
		deadMap[deadArr[i]] = true
	}
	// 初始值写入队列
	step := 0
	n0 := "0000"
	visited[n0] = true
	q.Push(n0)
	// 循环，条件：队列不为空
	for q.Length() != 0 {
		n := q.Length()
		// 当前队列中所有的点弹出，向外扩散，并将扩散的点写入队列
		for i := 0; i < n; i++ {
			current := q.Pop().(string)
			if current == target {
				return step
			}
			for j := 0; j < 4; j++ {
				up := plusOne(current, j)
				if !visited[up] && !deadMap[up] {
					visited[up] = true
					q.Push(up)
				}
			}
			for k := 0; k < 4; k++ {
				down := minusOne(current, k)
				if !visited[down] && !deadMap[down] {
					visited[down] = true
					q.Push(down)
				}
			}
		}
		// 步数加1
		step += 1
	}

	return -1
}

func plusOne(s string, i int) string {
	sBytes := []byte(s)
	if sBytes[i] == byte('9') {
		sBytes[i] = byte('0')
	} else {
		sBytes[i] += 1
	}
	return string(sBytes)
}

func minusOne(s string, i int) string {
	sBytes := []byte(s)
	if sBytes[i] == byte('0') {
		sBytes[i] = byte('9')
	} else {
		sBytes[i] -= 1
	}
	return string(sBytes)
}

func main() {
	// s := "1234"
	// sBytes := []byte(s)
	// fmt.Println(sBytes)

	// res := plusOne("1339", 1)
	// fmt.Println("res: ", res)

	deadArr := []string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}
	res := GetOpenLockMin("8888", deadArr) // -1
	fmt.Println(res)
}
