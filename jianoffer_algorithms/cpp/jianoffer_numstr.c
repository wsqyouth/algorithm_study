#include <iostream>
#include <vector>
#include <stack>
#include <queue>
#include <algorithm>
using namespace std;
/*
请实现一个函数用来判断字符串是否表示数值（包括整数和小数）。例如，字符串"+100","5e2","-123","3.1416"和"-1E-16"都表示数值。 但是"12e","1a3.14","1.2.3","+-5"和"12e+4.3"都不是。
*/

void isDigit(char *&pstr)
{
	// 判断数字
	while (*pstr != 0 && *pstr >= '0' && *pstr <= '9')
	{
		pstr++;
	}
}
bool isExponent(char *&pstr)
{
	if (*pstr != 'e' && *pstr != 'E')
		return false;
	++pstr;

	if (*pstr == '+' || *pstr == '-')
		++pstr;
	// 仅仅出现符号位
	if (*pstr == '\0')
		return false;

	// 判断数字
	isDigit(pstr);

	return (*pstr == '\0') ? true : false;
}

bool isNumeric(char *string)
{
	if (string == NULL)
		return false;
	char *pstr = string;

	if (*pstr == '+' || *pstr == '-')
		++pstr;

	if (*pstr == '\0')
		return false; // 仅仅出现符号位，非法数字

	// 判断数字
	isDigit(pstr);

	int isExpon = true;
	if (*pstr != '\0' && *pstr == '.')
	{
		++pstr;
		// 判断数字
		isDigit(pstr);

		// 判断指数
		if (*pstr == 'e' || *pstr == 'E')
		{
			isExpon = isExponent(pstr);
		}
	}
	// 判断指数
	else if (*pstr != '\0' && (*pstr == 'e' || *pstr == 'E'))
	{
		isExpon = isExponent(pstr);
	}
	else if (*pstr != '\0')
	{
		return false;
	}

	return isExpon && (*pstr == '\0');
}

int main()
{

	cout << isNumeric("12") << endl;
	cout << isNumeric("+-12") << endl;
	cout << isNumeric("-12") << endl;
	cout << isNumeric("+12.13") << endl;
	cout << isNumeric("+12.13E2") << endl;
	cout << isNumeric("+12.13E-2") << endl;
	cout << isNumeric("+12.13E+23") << endl;
	cout << isNumeric("12E+4.3") << endl;
	cout << isNumeric("1a3.14") << endl;
	cout << isNumeric("1.2.3") << endl;

	cout << isNumeric("") << endl;
	cout << isNumeric("\0") << endl;
	cout << isNumeric("+") << endl;
	cout << isNumeric("+13.") << endl;
	return 0;
}
