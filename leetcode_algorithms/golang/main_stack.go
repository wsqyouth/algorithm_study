package main

import (
	"fmt"
)

/*
func main() {
	isValidTest()
}
*/
func isValidTest() {
	str := "()[]{}"
	fmt.Println(isValid(str))
}

// lc20 判断括号有效性
func isValid(s string) bool {
	length := len(s)
	if length == 0 || length%2 == 1 {
		return false
	}
	stack := make([]byte, 0)
	for i := 0; i < length; i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else if len(stack) > 0 && leftOf(s[i]) == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			//和最近的左括号不匹配
			return false
		}
	}
	return len(stack) == 0
}

func leftOf(s byte) byte {
	signMap := map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}
	return signMap[s]
}

// lc 20  12.19 在lc上写一个使用switch case更简洁的版本.
func isValidNew(s string) bool {
	stack := []rune{}
	for _, ch := range s {
		switch ch {
		case '(', '{', '[':
			stack = append(stack, ch)
		case ')', '}', ']':
			if len(stack) == 0 {
				return false // 栈为空，无法匹配，返回false
			}
			top := stack[len(stack)-1]
			if (ch == ')' && top == '(') || (ch == '}' && top == '{') || (ch == ']' && top == '[') {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}

/*
栈的应用:
树的非递归遍历主要是利用栈实现的
*/
