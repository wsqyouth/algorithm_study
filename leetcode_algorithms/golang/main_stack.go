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
