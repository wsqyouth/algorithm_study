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
    // threeSum 时间复杂度n*n*nlog(n) = O(N3)
    vector<vector<int>> fourSum(vector<int> &nums, int target)
    {
        //数组排序
        sort(nums.begin(), nums.end());
        vector<vector<int>> res;
        // 穷举第一个数
        for (int i = 0; i < nums.size(); i++)
        {
            vector<vector<int>> tuples = threeSum(nums, i + 1, target - nums[i]);
            //如果存在则构成四元组
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

        return res;
    }

    // threeSum 前提是有序数组
    vector<vector<int>> threeSum(vector<int> &nums, int begin, int target)
    {
        vector<vector<int>> res;
        // 穷举第一个数
        for (int i = begin; i < nums.size(); i++)
        {
            vector<vector<int>> tuples = twoSum(nums, i + 1, target - nums[i]);
            //如果存在则构成三元组
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

        return res;
    }

    // twoSum 前提是有序数组
    vector<vector<int>> twoSum(vector<int> &nums, int begin, int target)
    {
        int lo = begin, hi = nums.size() - 1;
        vector<vector<int>> res;
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
        return res;
    }
};

int main()
{
    // Input: nums = [1,0,-1,0,-2,2], target = 0
    // Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
    //  vector<int> vec = {1,0,-1,0,-2,2};
    //     vector<vector<int>> resVec = fourSum(vec, 0);
    // Input: nums = [2,2,2,2,2], target = 8
    // Output: [[2,2,2,2]]
    vector<int> vec = {2, 2, 2, 2, 2};

    Solution s;
    vector<vector<int>> resVec = s.fourSum(vec, 8);
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
