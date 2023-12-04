// 该数据结构含有指向父节点的指针。
// 参考: https://blog.csdn.net/wtyvhreal/article/details/45565383

struct TreeLinkNode
{
	int val;
	struct TreeLinkNode *left;
	struct TreeLinkNode *right;
	struct TreeLinkNode *next;
	TreeLinkNode(int x) : val(x), left(NULL), right(NULL), next(NULL)
	{
	}
};

// 求当前节点的下一个节点
TreeLinkNode *GetNext(TreeLinkNode *pNode)
{
	if (pNode == NULL)
		return NULL;

	TreeLinkNode *pNext = NULL;

	// 右子树的最左子节点
	if (pNode->right != NULL)
	{
		TreeLinkNode *pCur = pNode->right;
		while (pCur->left != NULL)
		{
			pCur = pCur->left;
		}

		pNext = pCur;
	}
	// 父节点是结果的左子节点
	else if (pNode->next != NULL)
	{
		TreeLinkNode *pCurrent = pNode;
		TreeLinkNode *pParent = pNode->next;
		while (pParent != NULL && pParent->left != pCurrent)
		{
			pCurrent = pParent;
			pParent = pCurrent->next;
		}

		pNext = pParent;
	}

	return pNext;
}

// 求两个节点的最低公共节点
TreeLinkNode *GetLastParent(TreeLinkNode *pNode1, TreeLinkNode *pNode2)
{
	if (pNode1 == NULL || pNode2 == NULL)
		return NULL;

	TreeLinkNode *pLastParent = NULL; // the answer

	stack<TreeLinkNode *> path1;
	stack<TreeLinkNode *> path2;

	// 求出链表1(节点1->根节点)
	if (pNode1->next)
	{
		path1.push(pNode1);
		TreeLinkNode *pParent1 = pNode1->next;
		while (pParent1 && pParent1 != NULL)
		{
			path1.push(pParent1);
			pParent1 = pParent1->next;
		}
	}

	// 求出链表2(节点2->根节点)
	if (pNode2->next)
	{
		path2.push(pNode2);
		TreeLinkNode *pParent2 = pNode2->next;
		while (pParent2 && pParent2 != NULL)
		{
			path2.push(pParent2);
			pParent2 = pParent2->next;
		}
	}
	// 栈中依次弹出,求出最低公共节点
	while (path1.size() && path2.size() && path1.top() == path2.top())
	{
		pLastParent = path1.top();
		path1.pop();
		path2.pop();
	}

	return pLastParent;
}

/*
	  1
   2     5
3    4      6
*/
int main()
{
	TreeLinkNode a(1);
	TreeLinkNode b(2);
	TreeLinkNode c(5);
	TreeLinkNode d(3);
	TreeLinkNode e(4);
	TreeLinkNode f(9);
	TreeLinkNode g(6);

	a.left = &b;
	a.right = &c;
	b.left = &d;
	b.right = &e;
	// c.left = &f;
	c.right = &g;

	b.next = &a;
	c.next = &a;
	d.next = &b;
	e.next = &b;
	g.next = &c;

	cout << GetLastParent(&b, &b)->val << endl;
}
