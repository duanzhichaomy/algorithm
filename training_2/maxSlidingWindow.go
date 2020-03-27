/**
给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口 k 内的数字。滑动窗口每次只向右移动一位。

返回滑动窗口最大值。

示例:

输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
输出: [3,3,5,5,6,7]
解释:

  滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
注意：

你可以假设 k 总是有效的，1 ≤ k ≤ 输入数组的大小，且输入数组不为空。
*/

package training_2

import "container/list"

// 写得好的
func MaxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}
	window := list.New()
	var ret []int
	for i, v := range nums {
		if i >= k && window.Front().Value.(int) <= i-k {
			window.Remove(window.Front())
		}
		for window.Len() > 0 && nums[window.Back().Value.(int)] <= v {
			window.Remove(window.Back())
		}
		window.PushBack(i)
		if i >= k-1 {
			ret = append(ret, nums[window.Front().Value.(int)])
		}
	}
	return ret
}

// 自己写的
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nums
	}

	quene := nums[:k]
	maxnum := max(quene)
	result := []int{}
	result = append(result, maxnum)

	for i := k; i < len(nums); i++ {
		val := cal(quene, nums[i], &maxnum)
		result = append(result, val)
	}

	return result
}

func cal(quene []int, value int, maxnum *int) int {
	out := quene[0]
	quene = quene[1:]
	quene = append(quene, value)

	if value > *maxnum {
		*maxnum = value
	} else {
		if out == *maxnum {
			*maxnum = max(quene)
		}
	}

	return *maxnum
}

func max(value []int) int {
	max := value[0]

	for _, val := range value {
		if val > max {
			max = val
		}
	}

	return max
}
