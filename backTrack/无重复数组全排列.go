package backtrack

// leetcode 46

var res [][]int

func Permute(nums []int) [][]int {
	res = make([][]int, 0)
	choice := make([]int, 0)
	backtrack(nums, choice)
	return res
}

func backtrack(nums []int, choice []int) {
	// 如果符合条件，就加入结果集，并return
	if len(choice) == len(nums) {
		res = append(res, choice)
		return
	}
	for i := 0; i < len(nums); i++ {
		// 排除不合法的选择
		if inArr(nums[i], choice) {
			continue
		}
		// 做选择
		choice = append(choice, nums[i])
		// 进入下一决策树
		backtrack(nums, choice)
		// 撤销选择
		choiceTmp := make([]int, 0)
		for h := 0; h < len(choice); h++ {
			if choice[h] == nums[i] {
				continue
			}
			choiceTmp = append(choiceTmp, choice[h])
		}
		choice = choiceTmp
	}
	return
}

func inArr(a int, arr []int) bool {
	for j := 0; j < len(arr); j++ {
		if arr[j] == a {
			return true
		}
	}
	return false
}
