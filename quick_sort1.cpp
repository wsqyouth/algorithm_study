https://www.cnblogs.com/sdlwlxf/p/5131793.html
https://blog.csdn.net/superinzaghi747/article/details/52765555
https://www.cnblogs.com/zuilehongdou/p/6197716.html
http://haoyuanliu.github.io/2016/12/18/Partition%E7%AE%97%E6%B3%95%E5%89%96%E6%9E%90/
http://blog.jobbole.com/105219/
http://www.jobbole.com/members/xuelang/
https://www.cnblogs.com/sunniest/archive/2015/06/30/4610408.html

#include <stdio.h>
#include <vector>
#include <stack>
#include <iostream>
using namespace std;



void traverse(vector<int> vec)
{
	vector<int>::iterator it;
	for (it = vec.begin(); it != vec.end(); ++it)
		cout << *it << " ";
	cout << endl;
}

void swap(int &x, int &y)
{
	int t = x;
	x = y;
	y = t;

}
int partition(vector<int> &nums, int begin, int end)
{
	int pivot = nums[begin];//枢轴(也可以是在begin和end之间的随机数)
	// Last position where puts the no_larger element.
	//凡是小于pivot的全部放到数组左端,pos指向<枢轴值的最后一个,pos++指向不满足条件的（用于交换，将满足条件的换过来)
	int pos = begin;
	for (int i = begin + 1; i < end; ++i)
	{
		if (nums[i] < pivot)
		{
			pos++;
			if (i != pos) //避免自身交换
				swap(nums[pos], nums[i]);
		}

	}
	swap(nums[pos], nums[begin]);
	return pos;
}

/*
int Partition(vector<int> &nums, int begin, int end)
{
	int pivot = nums[begin];//枢轴(也可以是在begin和end之间的随机数)
	while (begin < end)
	{
		while (begin < end && nums[--end] >= pivot);
		nums[begin] = nums[end];//尾部找到小于pivot的值

		while (begin < end && nums[++begin] <= pivot);
		nums[end] = nums[begin];//头部找到大于pivot的值
	}

	nums[begin] = pivot;
	return begin;

}
*/
int Partition(vector<int>&arr, int begin, int end)
{
	int pivot = arr[begin];
	while (begin < end)
	{
		while (begin < end && arr[--end] >= pivot);
		arr[begin] = arr[end];
		while (begin < end && arr[++begin] <= pivot);
		arr[end] = arr[begin];
	}
	arr[begin] = pivot;
	return begin;
}

//确定一趟排序元素的位置
int identityTheIndex(vector<int> &nums, int low, int high)
{
	int key = nums[low];  //用字表的第一个记录作为枢轴
	while (low<high)
	{
		while (low<high&&nums[high] >= key){
			high--;
		}
		nums[low] = nums[high]; //将比第一个小的移到低端

		while (low<high&&nums[low] <= key){
			low++;
		}
		nums[high] = nums[low]; //将比第一个大的移到高端
	}
	nums[low] = key; //枢轴基准归位

	return low;

}


void quickSort(vector<int> &nums, int begin, int end)
{

	if (begin >= end) return;

	int index = identityTheIndex(nums, begin, end);
	if (index>begin)
		quickSort(nums, begin, index-1);
	if (index<end)
		quickSort(nums, index+1, end);
}

int main()
{


	vector<int> vec;
	vec.push_back(5);
	vec.push_back(9);
	vec.push_back(2);
	vec.push_back(1);
	vec.push_back(4);

	vec.push_back(7);
	vec.push_back(5);
	vec.push_back(8);
	vec.push_back(3);
	vec.push_back(0);

	traverse(vec);
	//cout << vec.size() << endl;
	quickSort(vec, 0, 10);
	//cout << Partition(vec, 0, 10) << endl;
	traverse(vec);
	return 0;
}
