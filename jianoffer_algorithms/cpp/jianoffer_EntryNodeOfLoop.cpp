一个链表中包含环，请找出该链表的环的入口结点。

方法1：
// s.insert(node).second这里在插入的同时也判断了插入是否成功，如果不成功表明set中已经有该元素了
ListNode *EntryNodeOfLoop(ListNode *pHead)
{
	if (pHead == NULL || pHead->next == NULL)
		return NULL; // 空或者1个节点
	if (pHead->next->next == NULL)
		return NULL; // a->b->null 有两个节点
	set<ListNode *> nodeSet;
	while (pHead)
	{
		if (nodeSet.insert(pHead).second) // 没找到则插入
		{
			nodeSet.insert(pHead);
			pHead = pHead->next;
		}
		else
		{
			return pHead;
		}
	}
	return NULL;
}

方法二：
ListNode *EntryNodeOfLoop(ListNode *pHead)
{
	if (pHead == NULL || pHead->next == NULL)
		return NULL; // 空或者1个节点
	if (pHead->next->next == NULL)
		return NULL; // a->b->null 有两个节点

	ListNode *pFast = pHead;
	ListNode *pSlow = pHead;
	// 快2 慢1
	while (pFast != NULL && pFast->next != NULL)
	{
		pFast = pFast->next->next;
		pSlow = pSlow->next;
		// 相遇时跳出
		if (pFast == pSlow)
		{
			break;
		}
	}
	if (pFast == NULL || pFast->next == NULL)
		return NULL;
	// 快指针保持不变，从相遇点开始跑(这里为了方便看懂，加一个临时变量)，慢指针从头开始跑
	ListNode *pMeet = pFast;
	pSlow = pHead;
	while (pMeet != NULL && pSlow != NULL)
	{
		if (pMeet == pSlow)
			return pMeet; // 环入口节点

		pMeet = pMeet->next;
		pSlow = pSlow->next;
	}
	return NULL;
}

-- -- -- -- -- -- --这里附加测试代码：
#include <stdio.h>

#include <set>
#include <algorithm>
#include <iostream>
				  using namespace std;

struct ListNode
{
	int val;
	struct ListNode *next;
	ListNode(int x) : val(x), next(NULL)
	{
	}
};

// 创建链表节点除了使用二级指针方法外，也可以采用返回值的办法。
ListNode *createListNode(int val)
{
	ListNode *pNode = new ListNode(val); // 利用了构造函数
	return pNode;
}
// 链接两个链表节点
void connectListNode(ListNode *pCur, ListNode *pNext)
{
	if (pCur == NULL)
	{
		cout << "error to connect" << endl;
		return;
	}
	pCur->next = pNext;
}
// 求链表的长度
int getListLength(ListNode *pHead)
{
	int nLength = 0;
	ListNode *pNode = pHead;
	while (pNode != NULL)
	{
		nLength++;
		pNode = pNode->next;
	}
	return nLength;
}
void printNode(ListNode *pListHead)
{
	if (pListHead == NULL)
		return;
	while (pListHead)
	{
		printf("%d ", pListHead->val);
		pListHead = pListHead->next;
	}
	printf("\n");
}

class Solution
{
public:
	void addToTail(ListNode **pListHead, int val)
	{
		ListNode *pNodeNew = new ListNode(val);

		if ((*pListHead) == NULL)
		{
			*pListHead = pNodeNew;
		}
		else
		{
			// pNode->next = (*pListHead)->next;
			ListNode *pNode = (*pListHead);
			while (pNode->next != NULL)
				pNode = pNode->next;

			pNode->next = pNodeNew;
		}
	}
	// s.insert(node).second这里在插入的同时也判断了插入是否成功，如果不成功表明set中已经有该元素了
	ListNode *EntryNodeOfLoop(ListNode *pHead)
	{
		if (pHead == NULL || pHead->next == NULL)
			return NULL; // 空或者1个节点
		if (pHead->next->next == NULL)
			return NULL; // a->b->null 有两个节点

		ListNode *pFast = pHead;
		ListNode *pSlow = pHead;
		// 快2 慢1
		while (pFast != NULL && pFast->next != NULL)
		{
			pFast = pFast->next->next;
			pSlow = pSlow->next;
			// 相遇时跳出
			if (pFast == pSlow)
			{
				break;
			}
		}
		if (pFast == NULL || pFast->next == NULL)
			return NULL;
		// 快指针保持不变，从相遇点开始跑(这里为了方便看懂，加一个临时变量)，慢指针从头开始跑
		ListNode *pMeet = pFast;
		pSlow = pHead;
		while (pMeet != NULL && pSlow != NULL)
		{
			if (pMeet == pSlow)
				return pMeet; // 环入口节点

			pMeet = pMeet->next;
			pSlow = pSlow->next;
		}
		return NULL;
	}
	ListNode *FindFirstCommonNode(ListNode *pHead1, ListNode *pHead2)
	{
		if (pHead1 == NULL || pHead2 == NULL)
			return NULL;
		int nLength1 = getListLength(pHead1);
		int nLength2 = getListLength(pHead2);
		// 较长链表先走若干步
		if (nLength1 > nLength2)
		{
			int nDiff = nLength1 - nLength2;
			while (pHead1 != NULL && nDiff > 0)
			{
				pHead1 = pHead1->next;
				nDiff--;
			}
		}
		else if (nLength1 < nLength2)
		{
			int nDiff = nLength2 - nLength1;
			while (pHead1 != NULL && nDiff > 0)
			{
				pHead2 = pHead2->next;
				nDiff--;
			}
		}
		// 两个链表同时遍历，找到的第一个相同节点就是他们的公共节点
		while (pHead1 != NULL && pHead1 != NULL)
		{
			if (pHead1 == pHead2)
			{
				return pHead1;
			}
			pHead1 = pHead1->next;
			pHead2 = pHead2->next;
		}

		return NULL;
	}
};

int main()
{
	Solution sovle;

	// 第一个公共结点在链表中间
	// 1 - 2 - 3 \  
	//            6 - 7
	//     4 - 5 /
	ListNode *pNode1 = createListNode(1);
	ListNode *pNode2 = createListNode(2);
	ListNode *pNode3 = createListNode(3);
	ListNode *pNode4 = createListNode(4);
	ListNode *pNode5 = createListNode(5);
	ListNode *pNode6 = createListNode(6);
	ListNode *pNode7 = createListNode(7);
	ListNode *pNode8 = createListNode(8);
	ListNode *pNode9 = createListNode(9);

	connectListNode(pNode1, pNode2);
	connectListNode(pNode2, pNode3);
	connectListNode(pNode3, pNode4);
	connectListNode(pNode4, pNode5);
	connectListNode(pNode5, pNode6);
	connectListNode(pNode6, pNode4);

	// printNode(pNode1);
	// printNode(pNode4);

	ListNode *head = sovle.EntryNodeOfLoop(pNode1);
	if (head != NULL)
	{
		cout << head->val << endl;
	}
	return 0;
}
