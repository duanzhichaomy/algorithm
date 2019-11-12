/**
给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。

示例 1:

输入: "(()"
输出: 2
解释: 最长有效括号子串为 "()"
示例 2:

输入: ")()())"
输出: 4
解释: 最长有效括号子串为 "()()"
*/

package training_2

func longestValidParentheses(s string) int {
	//定义返回值
	max := 0

	//判断长度
	if len(s) == 0 || len(s) == 1 {
		return max
	}

	//定义栈来保存（索引
	var stack []int

	//(索引
	start := 0

	//遍历
	for i := 0; i < len(s); i ++ {
		//判断是否为(
		if s[i] == '(' {
			//记录索引
			stack = append(stack,i)
		} else {
			//判断是否为空栈
			if len(stack) == 0 {
				start = i + 1
				continue
			}

			//出栈
			stack = stack[:len(stack) -1]

			//判断是否为空
			if len(stack) == 0 {
				max = maxInt(max,i - start + 1)
			} else {
				max = maxInt(max,i - stack[len(stack) - 1])
			}

		}

	}

	//返回
	return max
}

func maxInt(a,b int) int {
	if a > b {
		return  a
	} else {
		return b
	}
}

