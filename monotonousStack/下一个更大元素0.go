package monotonousstack

/*
单调栈模版的变形题目：
现在给你一个数组T，这个数组存放的是近几天的气温，算法需返回一个数组，
计算对于每一天，还要至少等多少天才能等到一个更暖和的气温；如果等不到那一天，填0。
如：输入T=[73,74,75,71,69,72,76,73]，算法应该返回[1,1,4,2,1,1,0,0]
*/
func nextGreaterElement0(t []int) []int {
	// 套单调栈模版，小改动
	res := make([]int, 0)
	stack := make([]int, 0) //索引入栈
	// 倒着入栈
	for i := len(t) - 1; i >= 0; i-- {
		// 比身高，去掉中间的矮个子
		for len(stack) != 0 && t[i] >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		// 返回结果
		if len(stack) != 0 {
			res = append(res, stack[len(stack)-1]-i)
		} else {
			res = append(res, 0)
		}
		// 最后入栈：参与身高比较
		stack = append(stack, i)

	}
	return res
}
