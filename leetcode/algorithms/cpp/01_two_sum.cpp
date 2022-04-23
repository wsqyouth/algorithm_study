#include <iostream>
#include <algorithm>
#include <vector>
#include <map>
using namespace std;

// declare
void printVec(vector<int> vec);
vector<int> twoSum(vector<int> &nums, int target);
vector<vector<int>> twoSumVal(vector<int> &nums, int target);

int main()
{
    vector<int> vec = {3, 2, 4};
    vector<int> res = twoSum(vec, 6);
    cout << res.size() << endl;
    printVec(res);

    cout << "two sum val:" << endl;
    vec = {1,3,1,2,2,3};
    vector<vector<int>> resVec = twoSumVal(vec, 4);
    for (int i = 0; i < resVec.size(); i++)
    {
        printVec(resVec[i]);
    }
}

// 找到对应的下标
vector<int> twoSum(vector<int> &nums, int target)
{
    vector<int> res;
    int len = nums.size();
    map<int, int> index_map;
    for (int i = 0; i < len; i++)
    {
        index_map.insert(pair<int, int>(nums[i], i));
    }
    std::map<int, int>::iterator iter;
    for (int i = 0; i < len; i++)
    {
        int other = target - nums[i];
        iter = index_map.find(other);
        //找到且不是自身
        if (iter != index_map.end() && iter->second != i)
        {
            res.push_back(i);
            res.push_back(iter->second);
            return res;
        }
    }
    //没找到，则设置默认值
    res.push_back(-1);
    res.push_back(-1);
    return res;
}

//找到对应的对[无序]
vector<vector<int>> twoSumVal(vector<int> &nums, int target)
{
    //数组排序
    sort(nums.begin(), nums.end());
    int lo = 0, hi = nums.size() - 1;
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