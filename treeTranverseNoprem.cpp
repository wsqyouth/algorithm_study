参考：https://blog.csdn.net/gogogo_sky/article/details/70145550

在我们了解了最简单的递归遍历以后，现在来说说用栈辅助，非递归实现二叉树的前序，中序，后序遍历打印 
一：非递归实现前序遍历打印 
1. 前序遍历，访问的顺序为根->左->右； 
2. 前序遍历的思想： 
（1）首先定义一个Node* 的指针cur; 用来遍历二叉树；然后定义一个栈，里面存储二叉树的结点； 
（2）利用while循环，一开始cur指向树的根结点，先把cur的值打印出来，然后入栈，然后将cur的值更新为他的左结点，继续循环；
     直到cur为NULL停止；这时cur指向最左结点的左结点，为NULL，此时表示最左结点和其左子树已经被打印； 
（3）接下来，需要打印最左结点的右子树；取出栈顶元素，将cur更新为栈顶元素的右结点；删除栈顶结点；

void _PreOrder(TreeNode* root)//前序遍历打印(根，左，右)
{
	TreeNode*  cur = root;
	stack<TreeNode*>  s;
	while (cur || !s.empty())
	{
		//每次将当前结点访问了。然后入栈
		//当while循环结束的时候，cur指向最左结点的左结点为NULL
		while (cur)
		{
			//访问了以后还保存，是为了下面从上往上打印左右结点的时候可以找到对应结点的左右节点
			cout << cur->val << " ";//打印
			s.push(cur);//入栈
			cur = cur->left;//更新
		}
		//while循环出来表示整个数的根结点和其左子树的根结点均已被打印；
		//现在要做的是从下往上一次打印左右节点；
		TreeNode* top = s.top();//取出最左结点；
		s.pop();//将最左结点出栈
		cur = top->right;
	}
	cout << endl;
}

二：非递归实现中序遍历打印 
1. 中序遍历，访问的顺序为左->根->右； 
2. 中序遍历的思想： 
（1）首先定义一个Node* 的指针cur; 用来遍历二叉树；然后定义一个栈，里面存储二叉树的结点； 
（2）利用while循环，一开始cur指向树的根结点，将cur入栈，然后将cur的值更新为他的左结点，继续循环；
     直到cur为NULL停止；这时cur指向最左结点的左结点，为NULL， 
（3）接下来，需要打印最左结点的左子树；因为左字树为空，所以直接打印他的根，取出栈顶元素，这就是最后一颗树的根，
     打印，删除栈顶结点,将cur更新为栈顶元素的右结点；
void _InOrder(TreeNode* root)//中序遍历打印（左，根，右）
{
	TreeNode* cur = root;
	stack<TreeNode*> s;
	while (cur || !s.empty())
	{
		//一直压栈到最左结点,while循环结束
		//表示当前结点及左子树结点均已经被访问过，
		//出循环的时候cur指得是最左结点的左子树，为NULL
		while (cur)
		{
			s.push(cur);
			cur = cur->left;
		}
		//开始打印左结点
		TreeNode* top = s.top();//取到最左结点
		cout << top->val << " ";//打印最左结点
		s.pop();//从栈中删除最左结点

		//最左结点的左孩子和根结点都已经被访问，接下来用子问题的方式访问右孩子
		cur = top->right;
	}

	cout << endl;
}


三：非递归实现后序遍历打印 
1. 后序遍历，访问的顺序为左->右->根； 
2. 后序遍历的思想： 
（1）首先定义一个栈和一个记录上一个被访问节点的变量，每遇到一个节点，就对其数据进行访问；然后将其入栈，
     将当前节点的左孩子赋给它，循环此过程，直到最左节点被访问并被压入栈中。 
（2）出循环后用临时变量保存栈顶元素，只有当前节点的右子树为空或者其右子树已经访问过时，才能对当前节点进行访问，同时将栈顶元素出栈 
（3）将临时变量的右孩子赋给当前节点，用子问题的方式去访问其右子树。
void _EndOrder(TreeNode* root)//后续遍历打印(左，右，根)
{
	TreeNode* cur = root;//当前结点
	TreeNode* prev = NULL;//上一次打印的结点
	stack<TreeNode*>  s;
	while (cur || !s.empty())//只要当前结点或者栈中还有元素，则该二叉树一定还没有打印完
	{
		while (cur)//一直找到最左结点
		{
			s.push(cur);
			cur = cur->left;
		}//while循环出来cur指向最左结点的做结点为NULL

		TreeNode* top = s.top();//取出最左结点

		//如果结点的右孩子为空，或者右孩子已经被打印，则可以打印本结点
		if (top->right == NULL || top->right == prev)
		{
			cout << top->val << " ";
			s.pop();
			prev = top;//将prev更新为已经打印过的结点
		}
		//如果结点的右孩子不为空，且还没有被访问，则将cur更新为右孩子，继续while循环，
		else
		{
			cur = top->right;
		}
	}

	cout << endl;
}

/*
      1
   2     5
3    4      6
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
	cout << "preoder thee" << endl;
	preorder_print(&a, layer);
	_PreOrder(&a);
	cout << "inorder thee" << endl;
	inorder_print(&a, layer);
	_InOrder(&a);
	cout << "postorder thee" << endl;
	postorder_print(&a, layer);/**/
	_EndOrder(&a);

}
