//自周六以来的三天，将二叉树的基本重新做了一遍，目标是做最低公共祖先的题。

#include <iostream>
#include <vector>
#include <queue>
#include <unordered_map>
using namespace std;

struct TreeNode {
	int val;
	TreeNode *left;
	TreeNode *right;
	TreeNode(int x) :val(x), left(NULL), right(NULL){}
};

void preorder_print(TreeNode *node, int layer)
{
	if (!node)
		return;
	//根据层数，打印相应数量的"--"
	for (int i = 0; i < layer; ++i)
		cout << "---";
	
	cout << "[" << node->val << "]\n";
	preorder_print(node->left, layer + 1); //遍历左子树
	preorder_print(node->right, layer + 1);//遍历右子树
}

void inorder_print(TreeNode *node, int layer)
{
	if (!node)
		return;
	inorder_print(node->left, layer + 1); //遍历左子树

	//根据层数，打印相应数量的"--"
	for (int i = 0; i < layer; ++i)
		cout << "---";

	cout << "[" << node->val << "]\n";

	inorder_print(node->right, layer + 1);//遍历右子树
}
void inorderTraversal(TreeNode *root,vector<int> &vec) {
	if (!root) return;
	if (root->left)
		inorderTraversal(root->left, vec);
	vec.push_back(root->val);
	if (root->right)
		inorderTraversal(root->right, vec);
}
vector<int> inorderTraversal(TreeNode *root) {
	vector<int> res; 
	if (!root) return res;
	inorderTraversal(root, res);

	return res;
}

void postorder_print(TreeNode *node, int layer)
{
	if (!node)
		return;
	postorder_print(node->left, layer + 1); //遍历左子树
	postorder_print(node->right, layer + 1);//遍历右子树

	//根据层数，打印相应数量的"--"
	for (int i = 0; i < layer; ++i)
		cout << "---";

	cout << "[" << node->val << "]\n";

}

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


vector<vector<int> > levelOrder(TreeNode *root) {
	vector<vector<int> > res;
	if (root == NULL) return res;

	queue<TreeNode *> qu;
	qu.push(root);
	while (!qu.empty())
	{
		int level = qu.size();
		vector<int> levelPath;
		for (int i = 0; i < level; ++i)
		{
			TreeNode * temp = qu.front();
			qu.pop();
			//cout << temp->val << " ";
			levelPath.push_back(temp->val);
			if (temp->left != NULL)
				qu.push(temp->left);
			if (temp->right != NULL)
				qu.push(temp->right);
		}
		//将每一层加入结果集res中
		res.push_back(levelPath);
	}
	return res;
}


vector<vector<int> > zigzagLevelOrder(TreeNode *root) {

	vector<vector<int> > res;
	if (root == NULL) return res;

	queue<TreeNode *> qu;
	qu.push(root);
	while (!qu.empty())
	{
		int level = qu.size();
		vector<int> levelPath;
		for (int i = 0; i < level; ++i)
		{
			TreeNode * temp = qu.front();
			qu.pop();
			//cout << temp->val << " ";
			levelPath.push_back(temp->val);
			if (temp->left != NULL)
				qu.push(temp->left);
			if (temp->right != NULL)
				qu.push(temp->right);
		}
		if (res.size() % 2 != 0)
		{
			reverse(levelPath.begin(), levelPath.end());
		}
		//将每一层加入结果集res中
		res.push_back(levelPath);
	}


	return res;
}

void pathSumCore(TreeNode *root, int sum, int &curSum, vector<int> &path_val, vector<vector<int> > &res) {
	if (root == NULL) return;

	path_val.push_back(root->val);
	curSum += root->val;
	if (root->left == NULL && root->right == NULL && curSum == sum)
	{
		#if 0
		for (vector<int>::iterator it = path_val.begin(); it != path_val.end(); ++it)
			cout << *it << " ";
		cout << endl;
		#endif
		//将满足条件的路径push到结果集合res中
		res.push_back(path_val);
	}

	pathSumCore(root->left, sum, curSum, path_val,res);
	pathSumCore(root->right, sum, curSum, path_val, res);

	//当前节点退回至父节点时，将该数据去除
	curSum -= root->val;
	path_val.pop_back(); 
}


vector<vector<int> > pathSum(TreeNode *root, int sum) {
	vector<vector<int> > res;
	if (root == NULL) return res;

	int curSum = 0;
	vector<int> path_val;
	pathSumCore(root, sum, curSum,path_val,res);

	return res;
}
bool isSameTree(TreeNode *p, TreeNode *q) {
	if (p == NULL && q == NULL) return true;

	if (p == NULL || q == NULL) return false;
	if (p->val != q->val) return false;

	//二者根值相等时
	return isSameTree(p->left, q->left) && isSameTree(p->right, q->right);
}


int TreeDepth(TreeNode *root)
{
	if (root == NULL) return 0;

	int left = TreeDepth(root->left);
	int right = TreeDepth(root->right);

	return (left > right) ? (left + 1) : (right + 1);
}

bool isBalanced(TreeNode *root) {
	if (root == NULL) return true;

	int nLeft = TreeDepth(root->left);
	int nRight = TreeDepth(root->right);
	int diff = nLeft - nRight;
	if (diff < -1 || diff>1)
		return false;


	return isBalanced(root->left) && isBalanced(root->right);

}



#if 0
bool isValidBST(TreeNode *root) {
	if (root == NULL) return true;


	int curVal = root->val;

	if (root->left != NULL && root->left->val > curVal)
		return false;
	if (root->right != NULL && root->right->val < curVal)
		return false;

	bool left = true;
	if (root->left)
		left = isValidBST(root->left);

	bool right = true;
	if (root->right)
		right = isValidBST(root->right);


	return left&&right;
}
#endif


void reConstructBinaryTree1(vector<int> pre, vector<int> vin) {
	if (pre.empty() || vin.empty()) return;
	int size = pre.size(); //等于vin.size()
	int index = 0;
	for (; index < size; ++index)
	{
		if (vin[index] == pre[0])
			break;
	}
	vector<int> pre_left, pre_right, vin_left, vin_right;
	for (int i = 0; i < index; ++i)
	{
		cout << pre[i + 1] << "pre left" << endl;
		pre_left.push_back(pre[i + 1]);
		cout << vin[i] << "inorder left" << endl;
		vin_left.push_back(vin[i]);
	}
	
	for (int j = index+1; j < size; ++j)
	{
		cout << pre[j] << "pre right" << endl;
		pre_right.push_back(pre[j]);
		cout << vin[j] << "inorder right" << endl;
		vin_right.push_back(vin[j]);
	}
	reConstructBinaryTree1(pre_left, vin_left);
	cout << "---------------"<<endl;
	reConstructBinaryTree1(pre_right, vin_right);
}

#if 0
TreeNode* reConstructBinaryTree(vector<int> pre, vector<int> vin) {
	if (pre.empty() || vin.empty()) return NULL;
	int size = pre.size(); //等于vin.size()
	int index = 0;
	for (; index < size; ++index)
	{
		if (vin[index] == pre[0])
			break;
	}
	vector<int> pre_left, pre_right, vin_left, vin_right;
	//找出前序、中序序列中的左子树
	for (int i = 0; i < index; ++i)
	{
		pre_left.push_back(pre[i + 1]);
		vin_left.push_back(vin[i]);
	}
	//找出前序、中序序列中的右子树
	for (int j = index + 1; j < size; ++j)
	{
		pre_right.push_back(pre[j]);
		vin_right.push_back(vin[j]);
	}
	//构造根节点
	TreeNode *root = new TreeNode(pre[0]);
	root->left = reConstructBinaryTree(pre_left, vin_left);
	root->right = reConstructBinaryTree(pre_right, vin_right);

	return root;
}
#endif

TreeNode* reConstructBinaryTree(vector<int> &pre, vector<int> &vin, int startPre, int endPre, int startInorder, int endInorder) {
	
	int rootVal = pre[startPre];
	//构造根节点
	TreeNode *root = new TreeNode(rootVal);
	if (endPre == startPre || endInorder == startInorder) return root;

	int index = startInorder;
	for (; index <= endInorder; ++index)
	{
		if (vin[index] == rootVal)
			break;
	}
	
	//左子树元素个数,右子树元素个数均根据中序遍历来找
	int leftLength = index - startInorder;
	int rightLength = endInorder - index;

	if (leftLength>0)
		root->left = reConstructBinaryTree(pre, vin, startPre + 1, startPre + leftLength, startInorder, index - 1);  //startPre + leftLength
	if (rightLength>0)
		root->right = reConstructBinaryTree(pre, vin, startPre + leftLength + 1, endPre, index + 1, endInorder);

	return root;
}

TreeNode* reConstructBinaryTree(vector<int> pre, vector<int> vin) {
	if (pre.empty() || vin.empty() || pre.size() != vin.size() ) return NULL;
	int size = pre.size(); //等于vin.size()

	return reConstructBinaryTree(pre, vin, 0, size - 1, 0, size - 1);
}

TreeNode *buildTree(vector<int> &inorder, vector<int> &postorder,int startInorder, int endInorder,int startPost,int endPost) {
	//从后续遍历中查找根节点
	int rootVal = postorder[endPost];
	//构造根节点
	TreeNode *root = new TreeNode(rootVal);
	if (endPost == startPost || endInorder == startInorder) return root;

	//在中序遍历序列查找index
	int index = endInorder;
	for (; index >= startInorder; --index)
	{
		if (inorder[index] == rootVal)
			break;
	}

	//左子树元素个数,右子树元素个数均根据中序遍历来找
	int leftLength = index - startInorder;
	int rightLength = endInorder - index;
	
	// 在根节点所在序列需要按照相对坐标来划分
	//左子树  后序[startPost，endPost-rightLength-1]
	if (leftLength>0)
		root->left = buildTree(inorder, postorder, startInorder, index - 1, startPost, endPost-rightLength-1);
	//右子树 后序[endPost-rightLength, endPost-1]
	if (rightLength>0)
		root->right = buildTree(inorder, postorder, index + 1, endInorder,endPost-rightLength, endPost-1);

	return root;
}
TreeNode *buildTree(vector<int> &inorder, vector<int> &postorder) {
	if (inorder.empty() || postorder.empty() || inorder.size() != postorder.size()) return NULL;
	int size = inorder.size(); //等于vin.size()

	return buildTree(inorder, postorder, 0, size - 1, 0, size - 1);
}
/*
      1
   2     5
3    4      6
*/
int main()
{
#if 0
	TreeNode a(8);
	TreeNode b(16);
	TreeNode c(10);
	TreeNode d(5);
	TreeNode e(7);
	TreeNode f(9);
	TreeNode g(11);

	a.left = &b;
	a.right = &c;
	b.left = &d;
	b.right = &e;
	c.left = &f;
	c.right = &g;

	int layer = 0;
	cout << "preoder thee" << endl;
	preorder_print(&a, layer);
	cout << "inorder thee" << endl;
	inorder_print(&a, layer);
	cout << "postorder thee" << endl;
	postorder_print(&a, layer);
	

	vector<vector<int> > res = levelOrder(&a);
	for (vector<vector<int>>::iterator ite = res.begin(); ite != res.end(); ite++)
	{
		vector<int> temp_vect = *ite;
		for (vector<int>::iterator itee = temp_vect.begin(); itee != temp_vect.end(); itee++)
			cout << *itee << " ";
		cout << endl;
	}
#endif
	//cout << isValidBST(&a) << endl;
	//前序遍历{1,2,4,7,3,5,6,8}和中序遍历序列{4,7,2,1,5,3,8,6}
	vector<int> pre;
	vector<int> vin;
	vector<int> post;
	pre.push_back(1);
	pre.push_back(2);
	pre.push_back(4);
	pre.push_back(7);
	pre.push_back(3);
	pre.push_back(5);
	pre.push_back(6);
	pre.push_back(8);

	vin.push_back(4);
	vin.push_back(7);
	vin.push_back(2);
	vin.push_back(1);
	vin.push_back(5);
	vin.push_back(3);
	vin.push_back(8);
	vin.push_back(6);
	
	post.push_back(7);
	post.push_back(4);
	post.push_back(2);
	post.push_back(5);
	post.push_back(8);
	post.push_back(6);
	post.push_back(3);
	post.push_back(1);

	TreeNode *root = buildTree(vin, post);

	//前序遍历
	int level = 0;
	preorder_print(root, level);
	cout << endl;
	
	inorder_print(root, level);
	cout << endl;

	vector<int> res1 = inorderTraversal(root);
	for (vector<int>::iterator it = res1.begin(); it != res1.end(); ++it)
		cout << *it << "__";
	cout << endl;

	//后续遍历
	postorder_print(root, level);
	cout << endl;

	vector<vector<int> > res = levelOrder(root);
	for (vector<vector<int>>::iterator ite = res.begin(); ite != res.end(); ite++)
	{
		vector<int> temp_vect = *ite;
		for (vector<int>::iterator itee = temp_vect.begin(); itee != temp_vect.end(); itee++)
			cout << *itee << " ";
		cout << endl;
	}
}
