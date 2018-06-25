题目一： 树的先序遍历、中序遍历、后序遍历、层次遍历
--------------------------------------------------
#include <iostream>
#include <vector>
#include <stack>
#include <queue>
#include <algorithm>
using namespace std;
struct TreeNode {
	int val;
	struct TreeNode *left;
	struct TreeNode *right;
	TreeNode(int x) :
		val(x), left(NULL), right(NULL) {
	}
};

void preOrder(TreeNode* pHead){
	//先序遍历，根 - 左 - 右
	if (pHead == NULL) return;
	cout << pHead->val << " ";
	preOrder(pHead->left);
	preOrder(pHead->right);
}

void inOrder(TreeNode* pHead){
	//中序遍历，左 - 根 - 右
	if (pHead == NULL) return;
	inOrder(pHead->left);
	cout << pHead->val << " ";//当前节点在其左右子节点之间访问， 左 - 根 - 右
	inOrder(pHead->right);
}

void posOrder(TreeNode* pHead){
	//后序遍历，左 - 右 - 根
	if (pHead == NULL) return;
	posOrder(pHead->left);
	posOrder(pHead->right);
	cout << pHead->val << " ";//当前节点在其左右子节点之后访问，左-右-根
}
//从上往下打印出二叉树的每个节点，同层节点从左至右打印。
void levelOrder(TreeNode *pHead) {

	if (pHead == NULL) return;

	deque<TreeNode *> dequeTreeNode;
	dequeTreeNode.push_back(pHead);
	while (!dequeTreeNode.empty())
	{
		TreeNode *temp = dequeTreeNode.front();
		//push val to vector
		cout << temp->val << " ";
		if (temp->left != NULL)
			dequeTreeNode.push_back(temp->left);
		if (temp->right != NULL)
			dequeTreeNode.push_back(temp->right);
		dequeTreeNode.pop_front();
	}

}

题目二： 树的深度
--------------------------------------------------
//利用层次遍历，每一层计数并处理一次
int TreeDepth(TreeNode *pRoot)
{
	if (!pRoot) return 0;
	queue<TreeNode *> qu;
	qu.push(pRoot);
	int depth = 0;
	while (!qu.empty())
	{
		int size = qu.size();
		depth++;
		for (int i = 0; i<size; i++)
		{
			TreeNode *temp = qu.front();
			cout << "__" << temp->val;
			qu.pop();

			if (temp->left != NULL)
				qu.push(temp->left);
			if (temp->right != NULL)
				qu.push(temp->right);
		}
		cout << endl;
	}
	return depth;
}


题目三： 二叉树先序遍历转换为双向链表
--------------------------------------------------
//对left,right判空，可以减少一次递归
//pLastInList使用指针引用的方法
void ConvertSub(TreeNode *pNode, TreeNode *&pLastInList)
{
	if (pNode == NULL) return;

	TreeNode *pCur = pNode;

	
	if (pCur->left != NULL)
		ConvertSub(pCur->left, pLastInList);

	//当前节点left指向pLastInList
	pCur->left = pLastInList;
	if (pLastInList != NULL)
		pLastInList->right = pCur; //尾结点right指向当前节点

	//更新尾结点
	pLastInList = pCur;

	//遍历右子树
	if (pCur->right != NULL)
		ConvertSub(pCur->right, pLastInList);

}

//null 1 3 6 8 10 null
TreeNode* Convert(TreeNode* pRootOfTree)
{
	TreeNode *pLastInList = NULL;
 	ConvertSub(pRootOfTree, pLastInList);

	//pLastInList指向双向链表的尾结点
	//我们根据该尾结点，一路向左，找到头结点
	TreeNode *pHeadInList = pLastInList;
	while (pHeadInList !=NULL && pHeadInList->left != NULL)
	{
		pHeadInList = pHeadInList->left;
	}
	return pHeadInList;
}



题目四： 二叉树求某特定和的路径(根节点到叶节点)
--------------------------------------------------
void FindPath(TreeNode * root, int target, vector<int> res, vector<vector<int> > &allres)
{
	if (root == NULL)  return;
	

	res.push_back(root->val);

	int isLeaf = root->left == NULL && root->right == NULL;
	if (isLeaf && target == root->val)
	{
	#if 0
		//打印路径点
		for (vector<int>::iterator iter = res.begin(); iter != res.end(); ++iter)
		{
			cout << *iter << " ";
		}
		cout << endl;
	#endif 
		allres.push_back(res);
	
	}
	//每遍历一个节点，更新target
	if (root->left != NULL)
		FindPath(root->left, target - root->val, res, allres);

	if (root->right != NULL)
		FindPath(root->right, target - root->val, res, allres);

	//返回父节点时，删除当前子节点
	res.pop_back();
}
vector<vector<int> > FindPath(TreeNode* root, int expectNumber) {
	
	vector<int> res;
	vector<vector<int> > allres;

	if (root == NULL)  return allres;
	
	FindPath(root, expectNumber, res, allres);

	return allres;
}



题目五： 二叉树判断相等
--------------------------------------------------
bool DoesTree1HasTree2(TreeNode* pRoot1, TreeNode* pRoot2)
{
	//递归比较至子树空
	if (pRoot2 == NULL)
		return true;

	if (pRoot1 == NULL)
		return false;

	if (pRoot1->val != pRoot2->val)
		return false;

	//二者都非空且值相等
	return  DoesTree1HasTree2(pRoot1->left, pRoot2->left) && DoesTree1HasTree2(pRoot1->right, pRoot2->right);
	
}
bool HasSubtree(TreeNode* pRoot1, TreeNode* pRoot2)
{
	//空树不是任意一个树的子结构
	if (pRoot2 == NULL) return false;

	if (pRoot1 == NULL) return false;

	bool res = false;
	if (pRoot1 != NULL && pRoot2 != NULL)
	{
		if (pRoot1->val == pRoot2->val)
		{
			res = DoesTree1HasTree2(pRoot1, pRoot2);
		}

		if (!res)
			res = HasSubtree(pRoot1->left, pRoot2);
		if (!res)
			res = HasSubtree(pRoot1->right, pRoot2);
	}

	return res;
}


题目六： 二叉树之字形打印
--------------------------------------------------
void print(TreeNode * pRoot)
{
	
	if (pRoot == NULL) return ;
	stack<TreeNode*> st0;
	stack<TreeNode*> st1;

	int depth = 1; //当前层数
	st0.push(pRoot);

	while (!st0.empty() || !st1.empty() )
	{
	
			
			//当前层数为奇数时，从左到右遍历子节点(栈的特性使下一次从右到左访问)
			while (!st0.empty())
			{ 
				TreeNode *curNode = st0.top();
				st0.pop();
				cout << curNode->val << "++";
				
				if (curNode->left != NULL)
					st1.push(curNode->left);
				if (curNode->right != NULL)
					st1.push(curNode->right);
			}
			
	
			//当前层数为偶数时，从右到左遍历子节点(栈的特性使下一次从左到右访问)
			while (!st1.empty())
			{
				TreeNode *curNode = st1.top();
				st1.pop();
				cout << curNode->val << "--";
				
				if (curNode->right != NULL)
					st0.push(curNode->right);
				if (curNode->left != NULL)
					st0.push(curNode->left);

			}
		
	}

}


题目七： 顺时针由外到里打印矩阵
--------------------------------------------------
void printMatrix(vector<vector<int> > matrix) {

	int rows = matrix.size();
	int columns = matrix[0].size();
	cout << rows << columns << endl;
	int start = 0;
	while (start*2 <rows && start*2<columns)
	{
		int endX = columns - 1 - start;
		int endY = rows - 1 - start;

		//打印行
		for (int i = start; i <= endX; ++i)
			cout << matrix[start][i] << " ";

		//打印列  至少两行(endY>start)
		if (endY > start)
		{
			for (int i = start + 1; i <= endY; ++i)
				cout << matrix[i][endX] << " ";
		}
		//打印行 至少两行(endY>start)  至少两列(endX>start)
		if (endY > start  && endX>start)
		{
			for (int i = endX - 1; i >= start;--i)
				cout << matrix[endY][i] << " ";
		}
		//打印列  至少三行(endY-1>start) 至少两列(endX>start)
		if (endY-1 > start  && endX>start)
		{
			for (int i = endY - 1; i > start; --i)
				cout << matrix[i][start] << " ";
		}
		start++;
	}
	cout << endl;
}

int main(){
	TreeNode a(1);
	TreeNode b(2);
	TreeNode c(3);
	TreeNode d(4);
	TreeNode e(5);
	TreeNode f(6);
	TreeNode g(7);

	a.left = &b;
	a.right = &c;
	b.left = &d;
	b.right = &e;
	c.left = &f;
	c.right = &g;

	TreeNode aa(8);
	TreeNode bb(8);
	TreeNode cc(9);
	aa.left = &bb;
	bb.left = &cc;
	
	vector<vector<int> > matrix;

	vector<int> v1{ 1, 2, 3, 4 }; 
	vector<int> v2{ 5,  6, 7, 8 }; 
	vector<int> v3{ 9, 10, 11, 12 };
	vector<int> v4{ 13, 14, 15,16};
	matrix.push_back(v1);
	matrix.push_back(v2);
	matrix.push_back(v3);
	matrix.push_back(v4);

	return 0;
}
