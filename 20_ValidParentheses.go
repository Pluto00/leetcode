package main

import "fmt"
/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
type stack struct {
	Length int
	Data   []int32
}

func (s *stack) Push(ch int32) {
	s.Data[s.Length] = ch
	s.Length++
}

func (s *stack) Pop() {
	s.Length--
}

func (s *stack) Top() int32 {
	return s.Data[s.Length-1]
}

func (s *stack) Empty() bool {
	return s.Length == 0
}

func isValid(s string) bool {
	strStack := stack{0, make([]int32, len(s))}
	for _, ch := range s {
		if ch == '(' || ch == '[' || ch == '{' {
			strStack.Push(ch)
		} else {
			if strStack.Empty() {
				return false
			}
			if ch == ')' && strStack.Top() != '(' {
					return false
			} else if ch == ']' && strStack.Top() != '[' {
					return false
			} else if ch == '}' && strStack.Top() != '{' {
					return false
			}
			strStack.Pop()
		}
	}
	return strStack.Empty()
}

func main() {
	fmt.Println(isValid("(]"))
}
