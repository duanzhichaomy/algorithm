/**
给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。

示例:

输入:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
输出: 7
解释: 因为路径 1→3→1→1→1 的总和最小。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-path-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
package training_7

func minPathSum(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	result := make([]int, cols)
	result[0] = grid[0][0]

	for j := 1; j < cols; j++ {
		result[j] = result[j - 1] + grid[0][j]
	}

	for i := 1; i < rows; i++ {
		result[0] = result[0] + grid[i][0]

		for j := 1; j < cols; j++ {
			result[j] = min(result[j], result[j - 1]) + grid[i][j]
		}
	}

	return result[cols - 1]
}