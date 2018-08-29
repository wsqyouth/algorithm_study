#include <iostream>
#include <vector>
#include <queue>
#include <string>
#include <iostream>
#include<algorithm>
#include<stack>
#include<sstream>
#include<functional>
using namespace std;

#pragma warning(disable:4996)

struct TreeNode {
	int val;
	TreeNode *left;
	TreeNode *right;
	TreeNode(int x) :val(x), left(NULL), right(NULL) {}
};


void level_print(TreeNode *node)
{
	if (!node)
		return;
	queue<TreeNode *> qu;
	qu.push(node);
	while (!qu.empty())
	{
		TreeNode * temp = qu.front();
		qu.pop();
		cout << temp->val << " ";
		if (temp->left != NULL)
			qu.push(temp->left);
		if (temp->right != NULL)
			qu.push(temp->right);
	}
	cout << endl;
}

char* Serialize(TreeNode *root) {
	if (root == NULL) return "#";
	string str = to_string(root->val);
	str.push_back(',');
	char *left = Serialize(root->left);
	char *right = Serialize(root->right);
	char *res = new char[strlen(left) + strlen(right) + str.size()];
	strcpy(res, str.c_str());//转换原字符串为C风格
	strcat(res, left);
	strcat(res, right);

	return res;
}

TreeNode* decode(char *&str) {
	if (*str == '#')
	{
		str++;
		return NULL;
	}
	int num = 0;
	while (*str !=',')
	{
		num = num * 10 + (*(str++) - '0');
	}
	str++;
	TreeNode *root = new TreeNode(num);
	root->left = decode(str);
	root->right = decode(str);

	return root;
}

TreeNode* Deserialize(char *str) {
	if (str == NULL) return NULL;
	return decode(str);
}


void _EndOrder(TreeNode* root)//后续遍历打印(左，右，根)
{
	if (root == NULL) return;
	TreeNode *cur = root;  //当前节点
	TreeNode *prev = NULL;//保存刚被访问的节点
	stack<TreeNode*> s;

	while (cur || !s.empty())
	{
		while (cur !=NULL)
		{
			s.push(cur);
			cur = cur->left;
		}

		TreeNode * top = s.top();
		if (top->right == NULL || top->right == prev)
		{
			cout << top->val << " ";
			s.pop();
			prev = top;
		}
		else
		{
			cur = top->right;
		}
	}
	cout << endl;
}

void _PreOrder(TreeNode *root) //前序遍历(根左右)
{
	if (root == NULL) return;
	TreeNode *cur = root;  //当前节点
	stack<TreeNode*> s;
	while (cur || !s.empty())
	{
		while (cur != NULL)
		{
			cout << cur->val << " ";
			s.push(cur);
			cur = cur->left;
		}

		TreeNode *top = s.top();
		s.pop();
		if (top->right != NULL)
		{
			cur = top->right;
		}
	}
	cout << endl;
}
void _InOrder(TreeNode *root) //中序遍历打印(左根右)
{
	if (root == NULL) return;
	TreeNode *cur = root;  //当前节点
	stack<TreeNode*> s;
	while (cur || !s.empty())
	{
		while (cur != NULL)
		{
			s.push(cur);
			cur = cur->left;
		}

		TreeNode *top = s.top();

		cout << top->val << " ";
		s.pop();

		if (top->right != NULL)
			cur = top->right;

	}
	cout << endl;
}
/*
	   1
   2     5
3    4  9   6
*/
int main()
{

	TreeNode a(1);
	TreeNode b(2);
	TreeNode c(5);
	TreeNode d(3);
	TreeNode e(4);
	TreeNode f(9);
	TreeNode g(6);

	a.left = &b;
	a.right = &c;
	b.left = &d;
	b.right = &e;
	c.left = &f;
	c.right = &g;

	int layer = 0;

	cout << "_PreOrder thee" << endl;
	_PreOrder(&a);


	cout << "_InOrder thee" << endl;
	_InOrder(&a);
	
	
	cout << "postorder thee" << endl;
	_EndOrder(&a);

	getchar();
	return 0;
}
