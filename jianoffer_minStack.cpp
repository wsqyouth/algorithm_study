//栈的使用：利用O(1)时间获取最小值

#include <stdio.h>
#include <vector>
#include <stack>
#include <iostream>
using namespace std;


/*
* 1.dataStack为存储数据的栈，minStack为存储最小值的栈；
* 2.push的时候将value值与minStack中的top值比较，小则minStack去push value，大则push top值
*/
void traverse(vector<int> vec)
{
	vector<int>::iterator it;
	for (it = vec.begin(); it != vec.end(); ++it)
		cout << *it << " ";
	cout << endl;
}

class Solution
{
public:
	void push(int value) {

		if (dataStack.empty() && minStack.empty())
		{
			dataStack.push(value);
			minStack.push(value);
		}
		else
		{
			int minVal=value;
			if (value > minStack.top())
			{
				minVal = minStack.top();
			}
			dataStack.push(value); 
			minStack.push(minVal); //将修改过的最小值push进去
		}
	}
	void pop() {
		dataStack.pop();
		minStack.pop();
	}
	int top() {
		return dataStack.top();
	}
	int min() {
		return minStack.top();
	}


	void traverseDataStack()
	{
		while (!dataStack.empty())
		{
			cout << dataStack.top() << " ";
			dataStack.pop();
		}
		cout << endl;
	}
	void traverseMinStack()
	{
		while (!minStack.empty())
		{
			cout << minStack.top() << " ";
			minStack.pop();
		}
		cout << endl;
	}
private:
	stack<int> dataStack; //数据栈
	stack<int> minStack;//最小数min栈
};


int main()
{


	Solution solve;
	solve.push(1);
	solve.push(-2);
	solve.push(3);
	solve.push(-5);
	solve.pop();
	solve.traverseDataStack();
	solve.traverseMinStack();


	return 0;
}
