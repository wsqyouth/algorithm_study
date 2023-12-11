### BM50 两数之和
```
/*
	for i := 0; i < len(numbers); i++ {
		anotherVal := target - numbers[i]
		for j := i + 1; j < len(numbers); j++ {
			if numbers[j] == anotherVal {
				return []int{i + 1, j + 1}
			}
		}
	}
*/
result := []int{}
valMap := make(map[int]int, len(numbers))
for i, v := range numbers {
	anotherVal := target - numbers[i]
	if val, exist := valMap[anotherVal]; exist {
		result = append(result, val+1)
		result = append(result, i+1) // 索引从1算
	}
	valMap[v] = i
}
```


总结: 好久没写make map定义竟然能搞错，这里要练习下纯手打代码，养成潜意识。
先从笨方法开始,之后使用哈希进行优化
这些题之前都做过，但是如何在牛客纯手工写代码不出错

### BM54 三数之和
思考: 看到这道题,第一种想法是确定一个值后,使用两数之和求剩余两个数，之后拼凑结果
第二种想法是双指针,需要维护几个指针满足条件时，进行拼凑，不过考虑到面试时间紧，没有那么多时间，仍然在第一种办法上优化

调试出错: 第一种办法是错误的
继续使用双指针
```
func threeSum(num []int) [][]int {
	result := [][]int{}
	// write code here
	sort.Ints(num)
	for i := 0; i < len(num); i++ {
		if i > 0 && num[i] == num[i-1] {
			continue
		}
		left, right := i+1, len(num)-1
		for left < right {
			sum := num[i] + num[left] + num[right]
			if sum == 0 {
				result = append(result, []int{num[i], num[left], num[right]})
				// 跳过重复元素
				for left < right && num[left] == num[left+1] {
					left++
				}
				for left < right && num[right] == num[right-1] {
					right--
				}
				left++
				right--
			} else if sum > 0 {
				right--
			} else if sum < 0 {
				left++
			}
		}
	}
	return result
}
```
扩展: 四数之和,接近三数之和怎么计算呢
```
对于四数之和的问题，可以使用类似于三数之和的思路进行解决。我们可以固定两个数，然后使用双指针来找到另外两个数，使得它们的和等于目标值。
对于接近三数之和的问题，可以使用类似的双指针方法，但需要记录最接近目标值的和。在遍历过程中，不断更新最接近目标值的和，并记录下来。最后返回最接近目标值的和即可。

sort.Ints(nums)
closestSum := math.MaxInt32
func abs(x int) int {
if x < 0 {
	return -x
}
return x
}
```
### BM62 斐波那契梳理
思路：使用make创建一个动态切片，把结果存进去,空间复杂度O(n)
```
func Fibonacci(n int) int {
	// write code here
	if n == 1 || n == 2 {
		return 1
	}
	dp := make([]int, n) // 注意避免range out
	dp[0] = 1
	dp[1] = 1
	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	// fmt.Println(dp)
	return dp[n-1]
}
```
### BM63 跳台阶
思路：直接使用两个变量存储，空间复杂度O(1)
```
func jumpFloor(number int) int {
	// write code here
	if number == 1 {
		return 1
	} else if number == 2 {
		return 2
	}
	var result int
	dp_i_2 := 1
	dp_i_1 := 2
	for i := 3; i <= number; i++ {
		result = dp_i_1 + dp_i_2
		dp_i_2 = dp_i_1
		dp_i_1 = result
	}
	return result
}
```
### BM9 删除链表的倒数第n个节点
思路: 快慢指针 不用slow,falst用错了一个地方
```
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// write code here
	dummyNode := &ListNode{}
	dummyNode.Next = head
	pSlowNode, pFastNode := dummyNode, head // 要小心pFastNode开始时执行的是头结点
	for i := 0; i < n; i++ {
		if pFastNode == nil {
			return nil
		}
		pFastNode = pFastNode.Next
	}

	for pFastNode != nil {
		pFastNode = pFastNode.Next
		pSlowNode = pSlowNode.Next
	}
	// 删除指向节点
	pSlowNode.Next = pSlowNode.Next.Next
	return dummyNode.Next
}
```

### BM4 合并两个有序的链表
### BM5 合并k个有序的链表


### BM94 接雨水问题
//     1.记录height中的每个元素，从左向右扫描并记录右边的最大高度
//     2.记录height中的每个元素，从右向左扫描并记录右边的最大高度
//     3.将左右位置元素对比取最小的元素，减去数组当前元素的高度
``` 
func maxWater(arr []int) int64 {
	// write code here
	n := len(arr)
	if n <= 0 {
		return 0
	}
	lMaxArray := make([]int, n) // 元素左侧最大值集合
	rMaxArray := make([]int, n) // 元素右侧最大值集合
	lMaxArray[0] = arr[0]
	rMaxArray[n-1] = arr[n-1]
	for i := 1; i < n; i++ {
		lMaxArray[i] = max(lMaxArray[i-1], arr[i])
	}
	for i := n - 2; i >= 0; i-- {
		rMaxArray[i] = max(rMaxArray[i+1], arr[i])
	}
	var res int64
	for i := 1; i < n-1; i++ {
		res += int64(min(lMaxArray[i],rMaxArray[i]) - arr[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {

	if a < b {
		return a
	}
	return b
}
```

### 位操作问题
1. 求一个数二进制中1的个数 n&(n-1)
2. 其他数出现2次，另一个数出现1次，异或操作
3. 其他数出现3次，只有一个数出现1次，笨办法: hashMap，稍微好办法: 数组去重后*3 - 数组求和= 待需求元素*2, 更好的办法: 继续异或求和,但是需要对结果区分
