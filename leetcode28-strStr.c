
// leetcode 28. Implement strStr()
// Problem Statement
// For a given source string and a target string, you should output the first
// index(from 0) of target string in source string.
// If target does not exist in source, just return -1.
// Example
// If source = "source" and target = "target", return -1.
// If source = "abcdabcdefg" and target = "bcd", return 1.
// Challenge
// O(n2) is acceptable. Can you implement an O(n) algorithm? (hint: KMP)

// -----------------------------------

#include <stdio.h>
#include <string.h>

int strStr(char *haystack, char *needle)
{

	int len = strlen(haystack);
	int len1 = strlen(needle);

	// printf("%d %d\n",len,len1);
	int i, j = 0;
	// 如果子串比母串长，返回-1
	if (len < len1)
	{
		return -1;
	}
	// 子串是空串
	if (len1 == 0)
	{
		return 0;
	}
	for (i = 0; i < len; i++)
	{
		// 母指针固定，子指针一直比较
		if (haystack[i] == needle[j])
		{

			while (haystack[i + j] == needle[j] && j < len1)
			{

				// 如果子指针到结尾，则查找到
				if (j == len1 - 1)
				{
					// printf("found it %d\n",j);
					return i; // 返回此时母指针的值，即子字符串首字母索引
				}
				j++;
			}
		}
		// 重新查找
		j = 0;
	}

	// 母指针到结尾，未查找到
	return -1;
}

int main()
{
	char str[20] = {"a"};
	char subStr[10] = {"a"};
	int res = strStr(str, subStr);
	printf("%d\n", res);
	return 0;
}
