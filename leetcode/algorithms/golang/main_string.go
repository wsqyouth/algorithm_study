package main

import (
	"fmt"
)

func main() {
	compareVersionTest()
}

func compareVersionTest() {
	fmt.Println(compareVersion("1.02", "1.001"))
}

// lc 165 比较版本号,使用双指针法的思想进行一一比较，不能使用for range循环遍历
func compareVersion(version1 string, version2 string) int {
	len1, len2 := len(version1), len(version2)
	var cur1, cur2 int
	for cur1 < len1 || cur2 < len2 {
		var num1, num2 int
		for ; cur1 < len1 && version1[cur1] != '.'; cur1++ {
			num1 = num1*10 + int(version1[cur1]-'0')
		}

		for ; cur2 < len2 && version2[cur2] != '.'; cur2++ {
			num2 = num2*10 + int(version2[cur2]-'0')
		}
		if num1 > num2 {
			return 1
		}
		if num1 < num2 {
			return -1
		}
		//跳过符号
		cur1++
		cur2++
	}

	return 0
}
