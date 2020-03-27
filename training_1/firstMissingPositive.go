/*
缺失的第一个正数

给定一个未排序的整数数组，找出其中没有出现的最小的正整数。

示例 1:
输入: [1,2,0]
输出: 3

示例 2:
输入: [3,4,-1,1]
输出: 2

示例 3:
输入: [7,8,9,11,12]
输出: 1
*/

package training_1

func firstMissingPositive(nums []int) int {
	numsLen := len(nums)
	result := numsLen + 1
	i := 0

	for i < numsLen {
		if nums[i] > 0 && (nums[i] <= numsLen) && (nums[i]-1 != i) && (nums[nums[i]-1] != nums[i]) {
			swap(nums, nums[i]-1, i)
		} else {
			i++
		}
	}

	for key, v := range nums {
		if (key + 1) == v {
			continue
		} else {
			result = key + 1
			break
		}
	}

	return result
}

func swap(data []int, i int, j int) {
	tmp := data[i]
	data[i] = data[j]
	data[j] = tmp
}

/************** 测试数据 *********************

s := []int{3,4,-1,1}
w := firstMissingPositive(s)

fmt.Println(w)

*********************************************/
