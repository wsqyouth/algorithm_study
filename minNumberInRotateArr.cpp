// 学习到知识点：
// 1) 旋转数组求最小值(二分法，首部和尾部和中间值相等时额外考虑)
// 2) linux环境下throw excption的用法

#include <iostream>
#include <stdlib.h>
#include <stdexcept>
#include <stack>
using namespace std;

int min(int arry[], int len) // 返回最小数的坐标
{
    if (arry == NULL || len <= 0)
        return -1;

    int start = 0;
    int end = len - 1;
    while (start < end)
    {
        // 如果首个元素小于最后一个元素，表明数组是排序的。
        if (arry[start] < arry[end])
            return start;

        // 当start指针和end指针隔壁的时候，返回end指针就是最小元素的坐标
        if (end - start == 1)
            return end;
        int mid = (start + end) / 2;
        // 如果arry[mid],arry[start]和arry[end]三个数相等，就只能使用顺序查找
        if (arry[mid] == arry[start] && arry[mid] == arry[end])
        {
            int index = start;
            for (int i = start + 1; i <= end; i++)
            {
                if (arry[i] < arry[index])
                    index = i;
            }
            return index;
        }

        // 如果中间元素小于末尾元素，那么表明中间元素在后半段数组中，修改end指针
        if (arry[mid] < arry[end])
        {
            end = mid;
        }
        // 如果中间元素大于首元素，那么表明中间元素在前半段数组中，修改start指针
        else if (arry[mid] > arry[start])
        {
            start = mid;
        }
    }
    return -1;
}

int minNum(int arry[], int len) // 返回最小数的值
{
    if (arry == NULL || len <= 0)
    {
        std::logic_error ex("非法输入");
        throw std::exception(ex);
    }

    int start = 0;
    int end = len - 1;
    while (arry[start] >= arry[end])
    {
        if (end - start == 1) // 如果start和end相差1则返回arry[end]
            return arry[end];

        int mid = (start + end) / 2;
        // 如果arry[mid],arry[start]和arry[end]三个数相等，就只能使用顺序查找
        if (arry[mid] == arry[start] && arry[mid] == arry[end])
        {
            int result = arry[start];
            for (int i = start + 1; i <= end; i++)
            {
                if (arry[i] < result)
                    result = arry[i];
            }
            return result;
        }

        if (arry[mid] >= arry[start]) // 如果中间元素大于首元素，则移动首指针
        {
            start = mid;
        }
        else if (arry[mid] <= arry[end])
        {
            end = mid;
        }
    }
    return arry[start]; // 如果一开始arry[start]<arry[end]表明数组是排序数组，返回arry[start]
}

int main()
{
    int arry[] = {3, 4, 5, 1, 2};
    int len = sizeof(arry) / sizeof(int);

    int index = min(arry, len);
    int minnum = minNum(arry, len);

    cout << "最小数在数组中的位置：" << index << endl;
    cout << "最小数的值：" << minnum << endl;

    return 0;
}
