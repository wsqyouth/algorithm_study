//判断弹出序列是否是入栈的弹出，是返回true,否则返回false
//核心思想见注释。这里难点在于何时插入、何时弹出

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

bool IsPopOrder(vector<int> pushV, vector<int> popV) {
	if (pushV.empty()) return false;
	
	stack<int> stackData;
	
	int i = 0; //pushV计数
	int j = 0;//popV计数
	while(i < pushV.size())
	{
		//队列头元素和栈顶不相等，则弹入队列一直往栈里push
		stackData.push(pushV[i++]);
		//队列头元素和栈顶相等，则二者一起向外“弹出”
		while (j<popV.size() && popV[j] == stackData.top())
		{
			stackData.pop();
			j++;
		}
	}
	

	return stackData.empty();
}
int main()
{
	vector<int> pushV;
	pushV.push_back(1);
	pushV.push_back(2);
	pushV.push_back(3);
	pushV.push_back(4);
	pushV.push_back(5);

	vector<int> popV;
	popV.push_back(4);
	popV.push_back(5);
	popV.push_back(3);
	popV.push_back(2);
	popV.push_back(1);
	
	cout << IsPopOrder(pushV, popV) << endl;


	return 0;
}
