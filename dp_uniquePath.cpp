leetcode原题：
从左上角走到右下角，路径条数。


无障碍版本：

int uniquePaths(int m, int n) {
	
	//二维初始化
	vector< vector<int> > dp(m, vector<int>(n, 0));

	//dp初始化 第一行全为1
	for (int j = 0; j < n; j++)
		dp[0][j] = 1;
	//dp初始化，第一列全为1
	for (int i = 0; i < m; i++)
		dp[i][0] = 1;

	for (int i = 1; i < m; ++i)
		for (int j = 1; j < n; ++j)
		{
			dp[i][j] = dp[i - 1][j] + dp[i][j - 1];
		}
	//返回bottom-right
	return dp[m - 1][n - 1];
}


--------------------
有障碍版本

#include <iostream>
#include <vector>
using namespace std;


void printVec(vector<vector<int> >  vec)
{
	vector<int> vec_tmp;

	for (vector<vector<int>>::iterator iter = vec.begin(); iter != vec.end(); iter++)
	{
		vec_tmp = *iter;
		for (vector<int>::iterator it = vec_tmp.begin(); it != vec_tmp.end(); it++)
			cout << *it << " ";
		cout << endl;
	}
}

//首先判断是否有障碍物，有则初始化为0，否则根据左临近值、上临近值判断(如果存在的话)
int uniquePathsWithObstacles(vector<vector<int> > &obstacleGrid) {

	int ranks = obstacleGrid.size();
	int columns = obstacleGrid[0].size();

	//二维初始化
	vector< vector<int> > dp(ranks, vector<int>(columns, 0));

	dp[0][0] = obstacleGrid[0][0] == 1 ? 0 : 1;

	//dp初始化 第一行由左临近值推导
	for (int j = 1; j < columns; ++j)
	{
		if (obstacleGrid[0][j] == 1)
		{
			dp[0][j] = 0;
		}
		else
		{
			dp[0][j] = dp[0][j - 1]; 
		}
	}

	//dp初始化 第一列由上临近值推导
	for (int i = 1; i < ranks; ++i)
	{
		if (obstacleGrid[i][0] == 1)
		{
			dp[i][0] = 0;
		}
		else
		{
			dp[i][0] = dp[i - 1][0];
		}
	}

	for (int i = 1; i < ranks; ++i)
		for (int j = 1; j < columns; ++j)
		{
			if (obstacleGrid[i][j] == 1)
			{
				dp[i][j] = 0;
			}
			else
			{
				dp[i][j] = dp[i - 1][j] + dp[i][j - 1];
			}
			
		}
	
	return dp[ranks - 1][columns - 1];
}

//简洁版本参考：
int uniquePathsWithObstacles1(vector<vector<int> > &a) {
	int i, j, m = a.size(), n = a[0].size();
	vector<vector<int> > dp(m, vector<int>(n, 0)); // 初始化成0
												   // 第一个格点的值与障碍数相反
	dp[0][0] = 1 - a[0][0];
	// 依次计算
	for (i = 0; i < m; ++i) {
		for (j = 0; j < n; ++j) {
			// 只有没有障碍才有通路
			if (a[i][j] == 0) {
				if (i == 0 && j != 0) dp[0][j] = dp[0][j - 1]; // 左
				else if (i != 0 && j == 0) dp[i][0] = dp[i - 1][0]; // 上
				else if (i != 0 && j != 0) dp[i][j] += dp[i - 1][j] + dp[i][j - 1]; // 左+上
			}
		}
	}
	return dp[m - 1][n - 1];
}
int main()
{

	vector<vector<int> > obstacleGrid(3,vector<int>(3,0));
	obstacleGrid[1][1] = 1; //中间有障碍物
	cout << uniquePathsWithObstacles1(obstacleGrid);
	//cout << uniquePaths(3, 4) << endl;

	getchar();
	return 0;
}
