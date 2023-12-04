// 一、树状三角形求最小路径和
void print2DVect(vector<vector<int>> &res)
{
	for (vector<vector<int>>::iterator ite = res.begin(); ite != res.end(); ite++)
	{
		vector<int> temp_vect = *ite;
		for (vector<int>::iterator itee = temp_vect.begin(); itee != temp_vect.end(); itee++)
			cout << *itee << " ";
		cout << endl;
	}
}

int minimumTotal(vector<vector<int>> &triangle)
{
	if (triangle.empty())
		return 0;

	// 初始化dp rows*rows 全为0
	int rows = triangle.size();
	vector<vector<int>> dp;
	for (int i = 0; i < rows; ++i)
	{
		dp.push_back(vector<int>(rows, 0)); // 每行push一个vector，rows个0
	}
	// 直接负责最后一行
	for (int j = 0; j < rows; ++j)
		dp[rows - 1][j] = triangle[rows - 1][j];

	// 由底向上,选取min值+当前Num值
	for (int i = rows - 2; i >= 0; --i)
	{
		for (int j = 0; j <= i; ++j)
			dp[i][j] = min(dp[i + 1][j], dp[i + 1][j + 1]) + triangle[i][j];
	}

	return dp[0][0];
}

int main()
{

	int test[][4] = {{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}};
	vector<vector<int>> tri;
	for (int i = 0; i < 4; ++i)
	{
		tri.push_back(vector<int>());
		for (int j = 0; j < i + 1; ++j)
		{
			tri[i].push_back(test[i][j]);
		}
	}

	cout << minimumTotal(tri) << endl;
}

// 二、矩阵中从左上到右下最小路径和
int minPathSum(vector<vector<int>> &grid)
{
	if (grid.empty())
		return 0;

	// print2DVect(grid);

	// 初始化dp rows*rows 全为0
	int rows = grid.size();
	int cols = grid[0].size();
	vector<vector<int>> dp;
	for (int i = 0; i < rows; ++i)
	{
		dp.push_back(vector<int>(cols, 0)); // 每行push一个vector，cols个0
	}

	// 计算第一行最小值
	dp[0][0] = grid[0][0];
	for (int j = 1; j < cols; ++j)
		dp[0][j] = dp[0][j - 1] + grid[0][j];

	// 首先计算每行第一列，之后根据左，上的最小值计算当前最小值
	for (int i = 1; i < rows; ++i)
	{
		dp[i][0] = dp[i - 1][0] + grid[i][0];
		for (int j = 1; j < cols; ++j)
			dp[i][j] = min(dp[i][j - 1], dp[i - 1][j]) + grid[i][j];
	}
	// print2DVect(dp);
	return dp[rows - 1][cols - 1];
}

int main()
{

	int test[][3] = {{1, 3, 1}, {1, 5, 1}, {4, 2, 1}};
	vector<vector<int>> tri;
	for (int i = 0; i < 3; ++i)
	{
		tri.push_back(vector<int>());
		for (int j = 0; j < 3; ++j)
		{
			tri[i].push_back(test[i][j]);
		}
	}

	cout << minPathSum(tri) << endl;
}
