

// 计数排序：
// http://www.cnblogs.com/hdk1993/p/4652814.html
// 计数排序的基本思想是：统计一个数序列中小于某个元素a的个数为n,则直接把该元素a放到第n+1个位置上。
// 当然当过有几个元素相同时要做适当的调整，因为不能把所有的元素放到同一个位置上。计数排序假设输入的元素都是0到k之间的整数。
#include <iostream>
using namespace std;

void CountSort(int a[], int b[], int array_size, int k)
{
    int *c = new int[k + 1];
    for (int i = 0; i <= k; i++)
    {
        c[i] = 0;
    }
    // c[i] contains the number of elements equal to i
    for (int i = 0; i < array_size; i++)
    {
        c[a[i]]++;
    }
    // c[i] contains the number of elements less than or equal to i
    for (int i = 1; i <= k; i++)
    {
        c[i] = c[i] + c[i - 1];
    }
    // confirm the num position
    for (int i = array_size - 1; i >= 0; i--)
    {
        b[c[a[i]] - 1] = a[i];
        c[a[i]]--;
    }
}

int main()
{
    int a[8] = {2, 5, 3, 0, 2, 3, 0, 3};
    int b[8];
    CountSort(a, b, 8, 5);
    for (int i = 0; i < 8; i++)
    {
        cout << b[i] << " ";
    }
    cout << endl;
    return 0;
}

/*
基数排序
基数排序是稳定性排序算法
http : // www.cnblogs.com/hdk1993/p/4662630.html
（1）MSD法实现 最高位优先法通常是一个递归的过程：<1>
    先根据最高位关键码K1排序，得到若干对象组，对象组中每个对象都有相同关键码K1。<2>
    再分别对每组中对象根据关键码K2进行排序，按K2值的不同，再分成若干个更小的子组，每个子组中的对象具有相同的K1和K2值。<3>
    依此重复，直到对关键码Kd完成排序为止。<4>
    最后，把所有子组中的对象依次连接起来，就得到一个有序的对象序列。
（2）LSD法实现
    最低位优先法首先依据最低位关键码Kd对所有对象进行一趟排序，
    再依据次低位关键码Kd -
    1对上一趟排序的结果再排序， 依次重复，直到依据关键码K1最后一趟排序完成，就可以得到一个有序的序列。
    使用这种排序方法对每一个关键码进行排序时，不需要再分组，而是整个对象组。
*/
#include <iostream>
using namespace std;

int getdigit(int x, int d) // get digit from right
{
    int a[] = {1, 1, 10, 100};
    return (x / a[d]) % 10;
}

void lsdradix_sort(int a[], int begin, int end, int d)
{
    int count[10];
    int *bucket = new int[end - begin + 1];
    for (int k = 1; k <= d; k++)
    {
        // init array of count
        for (int i = 0; i < 10; i++)
        {
            count[i] = 0;
        }
        // sum each digits numbers
        for (int i = begin; i <= end; i++)
        {
            count[getdigit(a[i], k)]++;
        }
        // ensure index in bucket (!!!pay attention to repeat numbers)
        for (int i = 1; i < 10; i++)
        {
            count[i] = count[i] + count[i - 1];
        }
        // fill bucket array
        for (int i = end; i >= begin; i--)
        {
            bucket[count[getdigit(a[i], k)] - 1] = a[i];
            count[getdigit(a[i], k)]--;
        }
        // gether the number from bucket
        for (int i = begin, j = 0; i <= end; i++, j++)
        {
            a[i] = bucket[j];
        }
    }
    delete[] bucket; //!!!
}

void msdradix_sort(int a[], int begin, int end, int d)
{
    int count[10];
    int *bucket = new int[end - begin + 1];
    // init array of count
    for (int i = 0; i < 10; i++)
    {
        count[i] = 0;
    }
    // sum each digits numbers
    for (int i = begin; i <= end; i++)
    {
        count[getdigit(a[i], d)]++;
    }
    // ensure index in bucket (!!!pay attention to repeat numbers)
    for (int i = 1; i < 10; i++)
    {
        count[i] = count[i] + count[i - 1];
    }
    // fill bucket array
    for (int i = end; i >= begin; i--)
    {
        bucket[count[getdigit(a[i], d)] - 1] = a[i];
        count[getdigit(a[i], d)]--;
    }
    // gether the number from bucket
    for (int i = begin, j = 0; i <= end; i++, j++)
    {
        a[i] = bucket[j];
    }
    delete[] bucket; //!!!
    // sort group data recursivily
    for (int i = 0; i < 10; i++)
    {
        int p1 = begin + count[i];
        int p2 = begin + count[i] - 1;
        if (p1 < p2 && d > 1)
        {
            msdradix_sort(a, p1, p2, d - 1);
        }
    }
}
int main()
{
    int array[] = {329, 457, 657, 839, 436, 720, 355};
    lsdradix_sort(array, 0, 6, 3);
    for (int i = 0; i < 7; i++)
    {
        cout << array[i] << " ";
    }
    cout << endl;
    int array1[] = {329, 457, 657, 839, 436, 720, 355};
    msdradix_sort(array1, 0, 6, 3);
    for (int i = 0; i < 7; i++)
    {
        cout << array1[i] << " ";
    }
    cout << endl;
    return 0;
}
