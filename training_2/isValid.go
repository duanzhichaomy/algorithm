/**
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true
示例 2:

输入: "()[]{}"
输出: true
示例 3:

输入: "(]"
输出: false
示例 4:

输入: "([)]"
输出: false
示例 5:

输入: "{[]}"
输出: true
**/

package training_2

func isValid(s string) bool {
	relation := map[byte]byte {
		'(' : ')',
		'[' : ']',
		'{' : '}',
	}

	if len(s) == 0 {
		return true
	}

	sByte := []byte(s)
	stack := []byte{}


	for _, v := range sByte {
		if _, ok := relation[v]; ok {
			stack = append(stack, v)
		} else {
			if len(stack) == 0 {
				return false
			}

			match := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]

			if relation[match] != v {
				return false
			}
		}
	}

	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}