/*
*基于堆排序算法，构建最大堆。时间复杂度为O(nlogk)
*如果用快速排序，时间复杂度为O(nlogn)；
*如果用冒泡排序，时间复杂度为O(n*k)
*/

我使用的堆排序(间接法使用最大堆，但是方法不太好)
----
vector<int> GetLeastNumbers_Solution(vector<int> input, int k) 
    {   
        if(input.empty()) return input;
        int length = input.size();
        
        if(k>length || k<=0 ) return vector<int>();
       
        vector<int> output;
        priority_queue<int> p; //默认创建大顶堆，队头元素最大
        for (int i = 0; i<length; ++i)
        {
            if (i < k)
                p.push(input[i]);//小于k个直接插入
            else
            {
                if (input[i] < p.top()) //待插入元素小于堆顶
                {
                    p.pop();
                    p.push(input[i]);
                }
            }
        }

        //遍历p并将元素放入容器vector
        while (!p.empty())
        {

            output.push_back(p.top());
            p.pop();
        }
        return output;
}

----
别人的最大堆(参考博客直接法)
https://blog.csdn.net/my_lovely_lemon_tree/article/details/78007316
------
#include <iostream>
#include <string>
#include <queue>
using namespace std;
vector<int> GetLeastNumbers_Solution(vector<int> input, int k) {
        int len=input.size();
        if(len<=0||k>len|| k<0) return vector<int>();
    
        vector<int> res(input.begin(),input.begin()+k);
        //建堆
        make_heap(res.begin(),res.end());
        
        for(int i=k;i<len;i++)
        {
            if(input[i]<res[0])
            {
                //先pop,然后在容器中删除
                pop_heap(res.begin(),res.end());
                res.pop_back();
                //先在容器中加入，再push
                res.push_back(input[i]);
                push_heap(res.begin(),res.end());
            }
        }
        //使其从小到大输出
        sort_heap(res.begin(),res.end());
         
        return res;
         
    }
int main()
{
	//数组初始化vector
	int iarray[] = { 4, 5, 1, 6, 2, 7, 3, 8 };
	//count： iarray数组个数
	size_t count = sizeof(iarray) / sizeof(int);
	//int数组初始化 ivec3
	vector<int> input(iarray, iarray + count);
	vector<int> output;
	output = GetLeastNumbers_Solution(input, 3);
	for (int i = 0; i < output.size();++i)
		cout << output[i] << endl;
	return 0;
}

----
基于快排Partiton思想 时间复杂度O(n)
	{
		int pivot = nums[begin];//第一个记录作为枢轴(也可是在begin和end之间的随机数)
		while (begin < end)
		{
			while (begin < end && nums[end] >= pivot)
			{
				end--;
			}
			nums[begin] = nums[end];//尾部找到小于pivot的值,移到低端
	
			while (begin < end && nums[begin] <= pivot)
			{
				begin++;
			}
			nums[end] = nums[begin];//头部找到大于pivot的值,移到高端
		}
		
		nums[begin] = pivot;//枢轴基准归位
	
		return begin;
	}

 
	vector<int> GetLeastNumbers_Solution(vector<int> input, int k) {
	 
	int len=input.size();
	if(len==0||k>len ||k<0) return vector<int>();
	if(len==k) return input;
	 
	int start=0;
	int end=len-1;
	int index=Partition(input,start,end);
	while(index!=(k-1))
	{
		if(index>k-1)
		{
			end=index-1;
			index=Partition(input,start,end);
		}
		else
		{
			start=index+1;
			index=Partition(input,start,end);
		}
	}
	 
	vector<int> res(input.begin(), input.begin() + k);
	 
	return res;
}
