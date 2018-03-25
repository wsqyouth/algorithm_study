不一定将操作放到pop里面，在push的时候也可以。
1，将主栈的元素全部push到副栈中
2，将新的元素push到副栈尾部
3，将副栈的元素全部push到主栈


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

class Solution
{
public:
	void push(int node) {
		//stack2为主(比如自顶向下为1,2,3)，stack1为temp(本轮循环结束时stack1自底向上为1,2,3)
		while (!stack2.empty())
		{
			stack1.push(stack2.top());//先把stack2的元素全部push到stack1中
			stack2.pop();
		}
		stack1.push(node); //放到stack1的栈顶(比如自底向上为1,2,3,4)
		while (!stack1.empty())
		{
			stack2.push(stack1.top());//把stack1的元素全部push到stack2中
			stack1.pop();
		}
		//此时stack2中自顶向下元素为1,2,3,4。顶上为最先进入的值，最先出去
	}

	int pop() {
		if (!stack2.empty())
		{
			int popVal = stack2.top();
			stack2.pop();
			return popVal;
		}
		else
		{
			return NULL;
		}
		
	}
	void traverseStack1()
	{
		while (!stack1.empty())
		{
			cout << stack1.top() << " ";
			stack1.pop();
		}
		cout << endl;
	}
	void traverseStack2()
	{
		while (!stack2.empty())
		{
			cout << stack2.top() << " ";
			stack2.pop();
		}
		cout << endl;
	}
private:
	stack<int> stack1;
	stack<int> stack2;
};


int main()
{
	

	Solution solve;
	solve.push(1);
	solve.push(2);
	solve.push(3);
	int res = solve.pop();
	solve.traverseStack2();


	return 0;
}

