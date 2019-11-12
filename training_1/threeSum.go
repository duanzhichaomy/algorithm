/*
三数之和

给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。
注意：答案中不可以包含重复的三元组。
*/

package training_1

import (
	"sort"
)

// 自己写的效率低的
func threeSum(nums []int) [][]int {
	var result [][]int
	var effect []int

	sort.Ints(nums)

	for a := 0; a < len(nums); a++ {
		if a > 0 && nums[a-1] == nums[a] { // 如果第二个数与前一个数一样，就跳过
			continue
		}

		for b := a + 1; b < len(nums); b++ {
			if (b > a+1) && nums[b-1] == nums[b] {
				continue
			}

			other := 0 - nums[a] - nums[b]
			for c := b + 1; c < len(nums); c++ {
				if (c > b+1) && nums[c-1] == nums[c] {
					continue
				}
				if nums[c] == other {
					effect = append(effect, nums[a], nums[b], nums[c])
					result = append(result, effect)
					effect = []int{}
				}
			}
		}
	}

	return result
}

// 别人写的效率高的
func threeSums(nums []int) [][]int {
	result := make([][]int, 0, len(nums)/3)
	numsHash := make(map[int]int)

	// 建立hash表
	for _, v := range nums {
		if count, ok := numsHash[v]; ok {
			numsHash[v] = count + 1
		} else {
			numsHash[v] = 1
		}
	}

	// 分类正负数
	neg := make([]int, 0, len(numsHash))
	pos := make([]int, 0, len(numsHash))
	for k := range numsHash {
		if k < 0 {
			neg = append(neg, k)
		} else {
			pos = append(pos, k)
		}
	}

	// 找出 0, 0, 0 的特殊情况
	if v, ok := numsHash[0]; ok && v >= 3 {
		result = append(result, []int{0, 0, 0})
	}

	// 求差, 查表
	for _, n := range neg {
		for _, p := range pos {
			diff := 0 - n - p
			if v, ok := numsHash[diff]; ok {
				if (diff == n || diff == p) && v >= 2 {
					result = append(result, []int{n, p, diff})
				}
				if diff < n || diff > p {
					result = append(result, []int{n, p, diff})
				}
			}
		}
	}

	return result
}


/************** 测试数据 *********************
nums := []int{13, 14, 1, 2, -11, -11, -1, 5, -1, -11, -9, -12, 5, -3, -7, -4, -12, -9, 8, -13, -8, 2, -6, 8, 11, 7, 7, -6, 8, -9, 0, 6, 13, -14, -15, 9, 12, -9, -9, -4, -4, -3, -9, -14, 9, -8, -11, 13, -10, 13, -15, -11, 0, -14, 5, -4, 0, -3, -3, -7, -4, 12, 14, -14, 5, 7, 10, -5, 13, -14, -2, -6, -9, 5, -12, 7, 4, -8, 5, 1, -10, -3, 5, 6, -9, -5, 9, 6, 0, 14, -15, 11, 11, 6, 4, -6, -10, -1, 4, -11, -8, -13, -10, -2, -1, -7, -9, 10, -7, 3, -4, -2, 8, -13}
ww := threeSum(nums)

fmt.Println(ww)
*********************************************/