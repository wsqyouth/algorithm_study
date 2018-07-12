#include <iostream>
#include <cstdio>
using namespace std;
//打印输出时，要符合一般习惯，把前面的0去掉,从左开始打印 
void PrintNum(char* numchar){  
    int i = 0;  
     
    while(numchar[i++] == '0');//找到数值从高位到低位第一个不为'0'的位置
    --i;  
    while ( numchar[i] != '\0' ){  
        cout<<numchar[i]; 
        i++;  
    }  
    cout<<endl;  
}  
void Print1ToMaxOfNDigitsRecursively(char* numchar, int length, int index)  
{  
    if ( index == length)  
    {  
        PrintNum(numchar);  
        return;  
    }  
    for ( int i = 0; i < 10; i++ )  
    {  
        numchar[index] = i +'0';  
        Print1ToMaxOfNDigitsRecursively(numchar, length, index+1);  
    }  

}  
void Print1ToMaxOfNDigits(int n)  
{  
    if ( n <= 0 )  
    {  
        cout<<n<<" is illegal"<<endl;  
        return;  
    }  
    char * numchar = new char[n+1];  
    numchar[n] = '\0';                  //先对字符串数组初始化  

    Print1ToMaxOfNDigitsRecursively(numchar,n,0); 
    
    delete[] numchar;  
}  

int main(){
   Print1ToMaxOfNDigits(3);

    getchar();
}
