/**
给定一个整数数组 nums ，找出一个序列中乘积最大的连续子序列（该序列至少包含一个数）。

示例 1:

输入: [2,3,-2,4]
输出: 6
解释: 子数组 [2,3] 有最大乘积 6。
示例 2:

输入: [-2,0,-1]
输出: 0
解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-product-subarray
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
package training_7

func maxProduct(nums []int) int {
	max := make([]int, len(nums))
	min := make([]int, len(nums))

	max[0] = nums[0]
	min[0] = nums[0]
	res := nums[0]

	for i := 1; i < len(nums); i++  {
		max[i] = toMax(toMax(max[i - 1] * nums[i], min[i - 1] * nums[i]), nums[i])
		min[i] = toMin(toMin(max[i - 1] * nums[i], min[i - 1] * nums[i]), nums[i])

		res = toMax(res, max[i])
	}

	return res
}