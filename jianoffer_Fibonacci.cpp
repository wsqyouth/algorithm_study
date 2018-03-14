递归法:
int Fibonacci(int n) {
            //递归必须注意终止条件
            if(n==0 ) return 0;
            else if(n==1) return 1;
            else return Fibonacci(n-1)+Fibonacci(n-2);         
     
        
    }
    
    
非递归法：
自下而上，循环实现，同时保持每一个阶段的中间变量
注：由于本题中函数返回值已规定为int，所以没必要设置为unsigned

#include <iostream>
#include <cstring>
using namespace std;

int Fibonacci(int n) {
    int lastLastValue,lastValue,value;
    lastLastValue = 0;
    lastValue = 1;
    if( n == 0)  return lastLastValue;
    else if(n == 1)  return lastValue;
    else {
        for(int i=2;i<=n;i++){
            value = lastLastValue + lastValue;
            lastLastValue = lastValue;// n-2
            lastValue = value;
        }
        return value;    
    }
        
}
	
int main() {
    
    for(int i=0;i<=10;i++)
        cout <<i<<"  " << Fibonacci(i)<<endl;
	return 0;
}
