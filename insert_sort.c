#include <stdio.h>
void insert_sort(int a[], int length)
{
	int t;
	int j;
	for (int i = 0; i < length; i++)
	{
		t = a[i]; // 将待排序的元素a[i]赋给t
		for (j = i - 1; j >= 0 && a[j] > t; j--)
		{
			a[j + 1] = a[j]; // 把比t大 （a[j]>t） 的元素，全部后移一个位置
		}
		a[j + 1] = t; // 把待排序的t插入到a[j+1].完成插入排序过程。
	}
}
int main()
{
	int a[5] = {12, 5, 41, 36, 89};
	insert_sort(a, 5);
	for (int i = 0; i < 5; i++)
		printf("%d  ", a[i]);

	return 0;
}

/*
直接插入排序的基本思想
	直接插入排序的基本思想是，把n个待排序的元素看成为一个有序表和一个无需表，开始时有序表中只包含一个元素，无序表中包含有n-1个元素。
排序过程中每次从无序表中取出第一个元素，将它插入到有序表中的适当位置，使之成为新的有序表，重复n-1次即可完成排序过程。
	把a[i]插入到a[0],a[1],...,a[i-1]之中的具体实施过程为：
	先把a[i]赋值给t，然后将t依次与a[i-1]，a[i-2]，....进行比较，将比t大的元素右移一个位置，
	直到发现某个j(0<=j<=i-1)，使得a[j]<=t或j为(-1),把t赋给a[j+1]。
*/