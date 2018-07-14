全文参考：https://blog.csdn.net/v_july_v/article/details/7041827


C语言版本：
----
#include <stdio.h>
#include <string.h>
int violentMatch(char *s,char *p)
{
    int sLen = strlen(s);
    int pLen = strlen(p);
    
    int i = 0;
    int j = 0;
    while(i<sLen && j<pLen)
    {
        if(s[i]==p[j])
        {
            i++;
            j++;
        }
        else
        {
            i = i-j+1;
            j = 0;
        }
    }
    
    //匹配
    if(j == pLen)
        return i-j;
    else 
        return -1;
}

void getNext(char *p,int next[])
{
    int pLen = strlen(p);
    
    int j = 0;
    int k = -1;
    next[0] = -1;
    while(j<pLen-1)
    {
        if(k==-1 || p[j] == p[k])
        {
            ++k;
            ++j;
            next[j] = k;
        }
        else
        {
            k = next[k];
        }
    }
}

void getNextVal(char *p,int next[])
{
    int pLen = strlen(p);
    
    int j = 0;
    int k = -1;
    next[0] = -1;
    while(j<pLen-1)
    {
        if(k==-1 || p[j] == p[k])
        {
            ++k;
            ++j;
            //next[j] = k;
            if(p[j] != p[k])
                next[j] = k;
            else
                next[j] = next[k];
        }
        else
        {
            k = next[k];
        }
    }
}

int kmpSearch(char *s,char *p)
{
    int sLen = strlen(s);
    int pLen = strlen(p);
    
    int next[pLen];
    //getNext(p,next);
    getNextVal(p,next);
    
    int i = 0;
    int j = 0;
    while(i<sLen && j<pLen)
    {
        if(j==-1 || s[i] ==p[j])
        {
            i++;
            j++;
        }
        else
        {
            j = next[j];
        }
    }
    //匹配
    if(j == pLen)
        return i-j;
    else 
        return -1;
}


int main(void) { 
	char *str="bacbababadababacambabacaddababacasdsd";
	char *substr="ababaca";
	printf("%s\n",str);
	printf("%s\n",substr);
	printf("%d\n",violentMatch(str,substr));
	printf("%d\n",kmpSearch(str,substr));
	
	return 0;
}



C++版本(KMP没有返回值，可修改)
------

#include <iostream>
#include <string>
#include <vector>
using namespace std;



int ViolentMatch(string &s, string &p)
{
	int sLen = s.size();
	int pLen = p.size();

	int i = 0;
	int j = 0;
	while (i < sLen && j < pLen)
	{
		if (s[i] == p[j])
		{
			//①如果当前字符匹配成功（即S[i] == P[j]），则i++，j++    
			i++;
			j++;
		}
		else
		{
			//②如果失配（即S[i]! = P[j]），令i = i - (j - 1)，j = 0    
			i = i - j + 1;
			j = 0;
		}
	}
	//匹配成功，返回模式串p在文本串s中的位置，否则返回-1
	if (j == pLen)
		return i - j;
	else
		return -1;
}

void getNext(const string &substr, vector<int> &next)
{
	next.clear();
	next.resize(substr.size());


	int j = 0;
	int k = -1;
	next[0] = -1;

	while (j<substr.size() - 1)
	{
		if (k == -1 || substr[j] == substr[k])
		{
			++j;
			++k;
			next[j] = k;

		}
		else
		{
			k = next[k];
		}
	}
}
void getNextVal(const string &substr, vector<int> &next)
{
	next.clear();
	next.resize(substr.size());


	int j = 0;
	int k = -1;
	next[0] = -1;

	while (j<substr.size() - 1)
	{
		if (k == -1 || substr[j] == substr[k])
		{
			++j;
			++k;
			if (substr[j] == substr[k])
			{
				next[j] = next[k];
			}
			else {
				next[j] = k;
			}
		}
		else
		{
			k = next[k];
		}
	}
}
void printVec(vector<int> & vec)
{
	for (vector<int>::iterator it = vec.begin(); it != vec.end(); ++it)
		cout << *it << " ";
	cout << endl;
}

void kmp(const string &str, const string &substr, vector<int> &next)
{
	int sLen = str.size();
	int pLen = substr.size();

	//getNext(substr, next);
	getNextVal(substr, next);

	int j = 0;
	int i = 0;
	while (i < sLen && j<pLen)
	{
		if (j == -1 || str[i] == substr[j])
		{
			i++;
			j++;
		}
		else
		{
			j = next[j];
		}
		/*
		if (j == pLen)
		{
			cout << "find in position at:" << i - j << endl;
			//假设查找多处，由于i已经移动到下一个位置了，无需考虑，仅将j=0重新查找即可
			j = 0;
		}
		*/

	}
	//仅仅查找第一个匹配的字符串
	if (j == pLen)
	{
		cout << "find in position at:" << i - j << endl;
	}
	

}
int main()
{
	string str = "abacABCababcabababab";
	string substr = "abab";
	//cin >> str >> substr;
	vector<int> next;
	kmp(str, substr, next);
	cout << ViolentMatch(str, substr);
	return 0;
}


//
bacbababadababacambabacaddababacasdsd
ababaca
