
/*
利用二维数组由上到下，由左到右递增的规律，
那么选取右上角或者左下角的元素a[row][col]与target进行比较，
当target小于元素a[row][col]时，那么target必定在元素a所在行的左边,
即col--；
当target大于元素a[row][col]时，那么target必定在元素a所在列的下边,
即row++；
*/

#include <iostream>
// #include <cstring>
// #include <cstdio>
#include <vector>
#include <algorithm>
using namespace std;

bool Find(int target, vector<vector<int>> array)
{

	int rows = array.size();	   // 行数
	int columns = array[0].size(); // 列数
	int row = 0;
	int column = columns - 1;

	while (row < rows && column >= 0)
	{
		if (array[row][column] == target)
		{
			return true;
			break;
		}
		else if (array[row][column] > target)
		{
			column--;
		}
		else
		{
			row++;
		}
	}
	return false;
}
int main()
{
	vector<vector<int>> array(4, vector<int>(4)); // 初始化二维数组行列
	array[0][0] = 1;
	array[0][1] = 2;
	array[0][2] = 8;
	array[0][3] = 9;
	array[1][0] = 2;
	array[1][1] = 4;
	array[1][2] = 9;
	array[1][3] = 12;

	array[2][0] = 4;
	array[2][1] = 7;
	array[2][2] = 10;
	array[2][3] = 13;

	array[3][0] = 6;
	array[3][1] = 8;
	array[3][2] = 11;
	array[3][3] = 15;

	for (int i = 0; i < array.size(); ++i)
	{
		for (vector<int>::iterator it = array[i].begin(); it != array[i].end(); it++)
			cout << *it << " ";
		cout << endl;
	}

	cout << Find(8, array) << endl;
}
