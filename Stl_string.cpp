C++ STL中最基本以及最常用的类或容器：
string 操作代码：
------
#include <iostream>
#include <string>

using namespace std;

int main()
{
    string s1;//初始化字符串，空字符串
    string s2 = s1; //拷贝初始化，深拷贝字符串
   
    string s3 = "I am Yasuo"; //直接初始化，s3存了字符串
    string s4(10, 'a'); //s4存的字符串是aaaaaaaaaa
    string s5(s4); //拷贝初始化，深拷贝字符串
    string s6("I am Ali"); //直接初始化
    string s7 = string(6, 'c'); //拷贝初始化，cccccc

    //string的各种操作
    string s8 = s3 + s6;//将两个字符串合并成一个
    s3 = s6;//用一个字符串来替代另一个字符串的对用元素

    cin >> s1;
    
    cout << s1 << endl;
    cout << s2 << endl; //s1初始化为空，深拷贝得s2为空 
    cout << s3 << endl;
    cout << s4 << endl;
    cout << s5 << endl;
    cout << s6 << endl;
    cout << s7 << endl;
    cout << s8 << endl;
    cout << "s7 size = " << s7.size() << endl; //字符串长度，不包括结束符
    cout << (s2.empty() ? "This string is empty" : "This string is not empty") << endl;;

    system("pause");
    return 0;
}

-------
#include <iostream>
#include <string>

using namespace std;

//C语言中用下标或者指针来访问数组元素
void printStrOld(string str)
{
	for (int i = 0; i < str.size(); i++)
	{
   		 cout << str[i] << endl;
	}
	
}

//迭代器:使用const_iterator使得访问元素时是能读不能写，这跟常量指针意思差不多
void printStrNew(string str)
{
	for (string::const_iterator it = str.begin(); it != str.end(); it++)
	{
 	   cout << *it << endl;
	}
	
}
int main()
{
   string s3 = "hello world"; //cin遇到空格停止，使用getline(cin,s3);获取行字符串比如"hello world" 
   
   printStrOld(s3);  
   cout << "-----"<<endl;
   printStrNew(s3);
     
  system("pause");
    return 0;
}

