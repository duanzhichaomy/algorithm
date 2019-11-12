/**
实现 int sqrt(int x) 函数。

计算并返回 x 的平方根，其中 x 是非负整数。

由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。

示例 1:

输入: 4
输出: 2
示例 2:

输入: 8
输出: 2
说明: 8 的平方根是 2.82842...,
     由于返回类型是整数，小数部分将被舍去。
*/

package training_3

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}

	low := 0
	high := x/2 + 1

	for low <= high {
		middle := (low + high) / 2

		if square(middle) > x{
			high = middle - 1
		} else if square(middle) == x {
			return middle
		} else if (square(middle) < x) && (square(middle + 1) > x) {
			return middle
		} else if square(middle) < x {
			low = middle + 1
		}
	}

	return 0
}

func square(x int) int {
	return x * x
}
