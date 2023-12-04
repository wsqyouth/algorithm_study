// leetcode344. Reverse String

// Write a function that takes a string as input and returns the string reversed.

// Example:
// Given s = "hello", return "olleh".

// -------

#include <stdio.h>
#include <string.h>

char *reverseString(char *s)
{
	int len = strlen(s);
	// printf("%s  %d\n",s,len);
	int i = 0;
	char temp;
	while (i < len / 2)
	{
		temp = s[len - 1 - i];
		s[len - 1 - i] = s[i];
		s[i] = temp;
		i++;
	}
	return s;
}
int main()
{
	char str[200] = {"hello1230"};

	printf("%s\n", reverseString(str));
	return 0;
}

-- --method2 : char *reverseString(char *moves)
{
	int len = strlen(moves);
	printf("%s %d\n", moves, len);
	char *start = moves;
	char *end = &moves[len - 1];

	char temp;
	while (start < end)
	{
		temp = *end;
		*end = *start;
		*start = temp;
		start++;
		end--;
	}

	return moves;
}
