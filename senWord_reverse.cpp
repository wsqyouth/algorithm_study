
-- -
// 经过反复尝试，终于打造出了属于自己的单词翻转程序。
// Input：I ma a .tneduts
// Output：I am a student.
// ----
// C语言版：

#include <stdio.h>
#include <string.h>
   // 逆转字符串
   void
   Reverse(char *pBegin, char *pEnd)
{
	if (pBegin == NULL || pEnd == NULL)
		return;
	while (pBegin < pEnd)
	{
		char temp = *pBegin;
		*pBegin = *pEnd;
		*pEnd = temp;
		pBegin++;
		pEnd--;
	}
}
// 逆转句子
char *ReverseSentence(char *str)
{
	if (str == NULL)
	{
		return str;
	}
	// 逆转整个句子
	Reverse(str, str + strlen(str) - 1);
	// 逆转各个单词
	char *pBegin = str;
	char *pEnd = str;
	while (*pEnd != '\0')
	{
		while (*pEnd != '\0' && *pEnd == ' ')
			++pEnd; // 过滤掉单词前边的空格，确定单词起始位置

		pBegin = pEnd; // 去除空格，使pBegin指向单词首字母

		while (*pEnd != '\0' && *pEnd != ' ')
			++pEnd; // 结束循环时，pEnd指向单词后第一个空格

		// 逆转单词
		Reverse(pBegin, pEnd - 1);
	}
	return str;
}

int main()
{

	char str4[] = "student. a am I";
	if (strcmp(ReverseSentence(str4), "I am a student.") == 0)
	{
		printf("reverse sentence ok\n");
	}
	return 0;
}

// -----
// C++语言版：
/**
参考剑指offer，第一次翻转句子，第二次翻转各个单词
end作为主指针，主遍历后寻找单词的首字母位置，单词后第一个空字符
reverse函数中参数2指向的是结束后的下一个字符
**/
#include <iostream>
#include <cstring>
#include <cstdio>
#include <algorithm>
using namespace std;

string ReverseSentence(string str)
{

	if (str.empty())
		return str;
	// 逆转整个句子
	reverse(str.begin(), str.end());
	cout << str << endl;
	// 逆转各个单词
	int start = 0;
	int end = 0;
	int length = str.length();
	//
	while (end < length)
	{
		while (end < length && str[end] == ' ')
			++end;	 // 过滤掉单词前边的空格，确定单词起始位置
		start = end; // start指向单词首字母

		while (end < length && str[end] != ' ')
			++end; // 结束循环时，end指向单词后第一个空格
		// 翻转单词
		reverse(str.begin() + start, str.begin() + end);
	}
	return str;
}

int main()
{
	string str;
	getline(cin, str);
	cout << ReverseSentence(str) << endl;

	return 0;
}
