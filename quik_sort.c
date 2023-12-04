// 快速排序是不稳定算法，这里测试的快排看起来非常简单，经常复习之
//  流程：一次划分后，确定枢轴位置
//       对于枢轴左右两边序列进行递归调用即可
//  时间复杂度： 最坏情况n^2(每次划分基准，都是当前无序序列去的最值),平均情况nlogn
//  空间复杂度： 来源于递归的栈空间消耗logn

// 选择基准：可以是序列的开始位置、中间都可以。调用 QuickSort(R,0,n-1)

// -------------------基准选择在序列的开始位置---------
/*参考数据结构p274(清华大学出版社，严蔚敏)*/
#include <iostream>
using namespace std;

#include <stdio.h>
// 确定一趟排序元素的位置
int identityTheIndex(int A[], int low, int high)
{

    int key = A[low]; // 用字表的第一个记录作为枢轴
    while (low < high)
    {
        while (low < high && A[high] >= key)
        {
            high--;
        }
        A[low] = A[high]; // 将比第一个小的移到低端

        while (low < high && A[low] <= key)
        {
            low++;
        }
        A[high] = A[low]; // 将比第一个大的移到高端
    }
    A[low] = key; // 枢轴基准归位

    return low;
}

// 递归进行排序
void quictSort(int A[], int low, int high)
{
    if (low < high)
    {
        int index = identityTheIndex(A, low, high);
        quictSort(A, low, index - 1);
        quictSort(A, index + 1, high);
    }
}
int main()
{

    int a[] = {57, 68, 59, 52, 72, 28, 96, 33, 24};
    quictSort(a, 0, sizeof(a) / sizeof(a[0]) - 1); /*这里原文第三个参数要减1否则内存越界*/
    for (int i = 0; i < sizeof(a) / sizeof(a[0]); i++)
    {
        cout << a[i] << " ";
    }

    return 0;
}
