/**
判断一个 9x9 的数独是否有效。只需要根据以下规则，验证已经填入的数字是否有效即可。

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-sudoku
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

package training_6

import (
	"strconv"
)

func IsValidSudoku(board [][]byte) bool {
	row := [9][9]bool{}
	column := [9][9]bool{}
	chunk := [9][9]bool{}

	for j:=0; j<9; j++ {
		for i:=0; i<9; i++ {
			numString := string(board[i][j])
			if numString == "." {
				continue
			}

			num, _ := strconv.Atoi(numString)
			num = num - 1

			block := i/3 * 3 + j / 3
			if row[j][num] || column[i][num] || chunk[block][num] {
				return false
			} else {
				row[j][num] = true
				column[i][num] = true
				chunk[block][num] = true
			}
		}
	}

	return false
}


