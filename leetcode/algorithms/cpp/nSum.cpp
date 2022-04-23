#include <iostream>
#include <vector>
#include <map>
#include <algorithm>
using namespace std;

// declare
void printVec(vector<int> vec);

class Solution
{
public:
    vector<vector<int>> fourSum(vector<int> &nums, int target)
    {
        //数组排序
        sort(nums.begin(), nums.end());
        return nSumTarget(nums, 4, 0, target);
    }
    // nSumTarget前提是有序数组,调用前一定要先排序
    vector<vector<int>> nSumTarget(vector<int> &nums, int n, int begin, int target)
    {
        vector<vector<int>> res;
        // 至少是2Sum,且数组元素个数不应该小于n
        if (n < 2 || nums.size() < n)
        {
            return res;
        }
        //2Sum是base case
        if (n == 2)
        {
            int lo = begin, hi = nums.size() - 1;
            while (lo < hi)
            {
                int sum = nums[lo] + nums[hi];
                //记录最初的值
                int left = nums[lo], right = nums[hi];
                if (sum < target)
                    while (lo < hi && nums[lo] == left)
                    {
                        lo++;
                    }
                else if (sum > target)
                    while (lo < hi && nums[hi] == right)
                    {
                        hi--;
                    }
                else
                {
                    res.push_back({left, right});
                    //去重
                    while (lo < hi && nums[lo] == left)
                    {
                        lo++;
                    }
                    while (lo < hi && nums[hi] == right)
                    {
                        hi--;
                    }
                }
            }
        }
        else
        {
            //n>2是递归运算
            // 穷举第一个数
            for (int i = begin; i < nums.size(); i++)
            {
                vector<vector<int>> tuples = nSumTarget(nums, n - 1, i + 1, target - nums[i]);
                //如果存在则构成n-1元组
                for (vector<int> &tuple : tuples)
                {
                    tuple.push_back(nums[i]);
                    res.push_back(tuple);
                }
                //跳过第一个数的重复情况
                while (i < nums.size() - 1 && nums[i] == nums[i + 1])
                {
                    i++;
                }
            }
        }

        return res;
    }
};

int main()
{
    // Input: nums = [1,0,-1,0,-2,2], target = 0
    // Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
    vector<int> vec = {1, 0, -1, 0, -2, 2};

    Solution s;
    vector<vector<int>> resVec = s.fourSum(vec, 0);
    cout << resVec.size() << endl;
    cout << "four sum val:" << endl;

    for (int i = 0; i < resVec.size(); i++)
    {
        printVec(resVec[i]);
    }
}

void printVec(vector<int> vec)
{
    vector<int>::iterator v = vec.begin();
    while (v != vec.end())
    {
        cout << " " << *v << " ";
        v++;
    }
    cout << " " << endl;
}
