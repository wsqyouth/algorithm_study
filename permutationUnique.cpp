// 全排列问题
//  情况1：数组不重复
//  情况2：数组中含有重复数字

#include <iostream>
#include <string>
#include <vector>
#include <algorithm>
#include <sstream>
#include <unistd.h>
#include <stdio.h>
#include <sys/types.h>
using namespace std;

class Solution
{
public:
	//[start,end]   num[end]
	bool isSwap(vector<int> num, int start, int end)
	{
		for (int i = start; i < end; ++i)
		{
			if (num[i] == num[end])
				return false;
		}
		return true;
	}
	void permutation(vector<vector<int>> &res, vector<int> &num, int begin)
	{

		if (begin == num.size() - 1)
		{
			res.push_back(num);
		}
		else
		{

			for (int i = begin; i < num.size(); ++i)
			{
				if (isSwap(num, begin, i))
				{
					swap(num[i], num[begin]);
					permutation(res, num, begin + 1);
					swap(num[i], num[begin]);
				}
			}
		}
	}
	vector<vector<int>> permute(vector<int> &num)
	{
		vector<vector<int>> res;
		if (num.empty())
			return res;

		permutation(res, num, 0);
		return res;
	}
};
int main(int argc, const char *argv[])
{

	Solution s;
	int arr_tmp[] = {1, 2, 1};
	vector<int> vec(arr_tmp, arr_tmp + 3);
	vector<vector<int>> ves = s.permute(vec);
	for (vector<vector<int>>::iterator it = ves.begin(); it != ves.end(); ++it)
	{
		for (int i = 0; i < (*it).size(); ++i)
			cout << (*it)[i] << " ";
		cout << endl;
	}

	return 0;
}
