package bfs

import (
	"strings"
)

var neighbors = [6][]int{{1, 3}, {0, 2, 4}, {1, 5}, {0, 4}, {1, 3, 5}, {2, 4}}

func SlidingPuzzle(board [][]int) int {
	const target = "123450"
	// 二维数组--》string
	s := make([]byte, 0, 6)
	for _, r := range board {
		for _, v := range r {
			s = append(s, '0'+byte(v))
		}
	}
	start := string(s)
	// 队列
	type pair struct {
		value string
		step  int
	}
	q := []pair{{start, 0}}
	// visited
	visited := map[string]bool{start: true}
	// 扩散子方法
	get := func(current string) (res []string) {
		currentBytes := []byte(current)
		pos0 := strings.Index(current, "0")
		for _, v := range neighbors[pos0] {
			currentBytes[pos0], currentBytes[v] = currentBytes[v], currentBytes[pos0]
			res = append(res, string(currentBytes))
			currentBytes[pos0], currentBytes[v] = currentBytes[v], currentBytes[pos0]
		}
		return
	}
	// 循环，条件：队列不为空
	for len(q) != 0 {
		current := q[0]
		q = q[1:]
		if current.value == target {
			return current.step
		}
		// 当前队列中所有的点向外扩散，且扩散的点写入队列
		for _, item := range get(current.value) {
			if !visited[item] {
				//step+=1
				q = append(q, pair{value: item, step: current.step + 1})
				visited[item] = true
			}
		}
	}

	return -1
}
