package binarysearch

// leetcode 33

func Search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] {
			// 说明左半部分是有序的
			if target >= nums[0] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}

		} else {
			// 说明右半部分是有序的
			if nums[mid] < target && target <= nums[len(nums)-1] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

// 笨办法，测试用例（175/198）

// func search(nums []int, target int) int {
//     if len(nums)<=0{
//         return -1
//     }
//     if len(nums)==1{
//         if nums[0]==target{
//             return 0
//         }
//         return -1
//     }
//     // 先找到k
//     left:=0
//     right:=len(nums)-1
//     k:=0
//     for left<=right{
//         if nums[left]<=nums[right]{
//             k = left
//             break
//         }
//         left+=1
//     }
//     if nums[0]<=target{
//         fmt.Println("k: ", k)
//         return test(nums, target, 0, k-1)
//     }
//     return test(nums,target, k, len(nums)-1)
// }

// func test(nums []int, target, leftIndex, rightIndex int)int{
//     // 先默写模版
//     left:=leftIndex
//     right:=rightIndex
//     for left<=right{
//         mid:=left+(right-left)/2
//         if nums[mid]<target{
//             left = mid+1
//         }else if nums[mid]>target{
//             right = mid-1
//         }else if nums[mid]==target{
//             return mid
//         }
//     }
//     return -1
// }
