#include <stdio.h>
#include <stdlib.h>
#include <stdarg.h>
void fun(char*fmt,...)
{
     int m;
     double d;
     char *ptr;
     va_list ap; //定义一个va_list类型变量
     va_start(ap,fmt);  //获取第二个参数的地址
     m = va_arg(ap,int);  //第二个参数是int类型，获取值
     d = va_arg(ap,double); //第三个参数是double类型，获取值
     ptr = va_arg(ap,char*); //第四个参数是char*类型，获取值
     va_end(ap);
     printf("%d\n",m);
     printf("%lf\n",d);
     printf("%s\n",ptr);
}
int main()
{
    fun("%d %f %s\n", 4, 5.4, "hello world");
    system("pause");
    return 0;
}
