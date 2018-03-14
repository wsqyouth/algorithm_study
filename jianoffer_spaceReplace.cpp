#include <iostream>
#include <cstring>
using namespace std;

void replaceSpace(char *str,int length) {
    char *pOld = str;
    int spaceCount = 0;
    //1.空格计数
    while(*pOld)
    {
        if(*pOld == ' ')
            spaceCount++;
        pOld++;
    }
    cout << spaceCount <<*(pOld)<<endl;
    //2.计算新的数组空间，并令pNew指向新数组结尾\0
    char *pNew = pOld + 2*spaceCount;
    //3.从后向前赋值，当pOld遇到空格时，pOld--,而pNew需要赋值替换的字符串并自减，直到二者相等
    while(pOld != pNew)
    {
        if(*pOld != ' ')
        {
            *(pNew--) = *(pOld--);
        }
        else {
            pOld--;
            *pNew-- = '0';
            *pNew-- = '2';
            *pNew-- = '%';
        }
    }
    cout << str<<endl;
}
	
int main() {
    char str[] = "We Are Happy";
    int len = strlen(str);
    replaceSpace(str,len);
    cout << str<<"  " <<len<<endl;
	return 0;
}
