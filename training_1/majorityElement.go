/*
求众数

给定一个大小为 n 的数组，找到其中的众数。众数是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
你可以假设数组是非空的，并且给定的数组总是存在众数。

示例 1:

输入: [3,2,3]
输出: 3
示例 2:

输入: [2, 2, 1, 1, 1, 2, 2]
输出: 2
 */
package training_1

import "fmt"

// 自己写的
func majorityElement(nums []int) int {
	majorityLen := len(nums) / 2
	hashMap := make(map[int]int)
	result := nums[0]

	for i := 0; i < len(nums); i++ {
		if count, ok := hashMap[nums[i]]; ok {
			hashMap[nums[i]] = count + 1

			if hashMap[nums[i]] > majorityLen {
				result = nums[i]
			}
		} else {
			hashMap[nums[i]] = 1
		}
	}

	return result
}

// 别人的
func majorityElements(nums []int) int {
	all := make(map[int]int)
	max := 0
	maxNum := 0
	for _, b := range nums {
		if all[b] != 0 {
			all[b] = all[b] + 1
		} else {
			all[b] = 1
		}
		if all[b] > maxNum {
			max = b
			maxNum = all[b]
		}
	}
	return max
}

func main() {
	s := []int{1}
	w := majorityElement(s)

	fmt.Println(w)
}
