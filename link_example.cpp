// 这个是之前自己在做链表基本操作时，写下的模板。结合牛客上推荐的数据结构，完成自己的功能。
// 主要包括：
// 链表合并、链表翻转，尾插法构建链表（创建时采用了二级指针的写法），打印链表功能。

#include <stdio.h>
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
class Solution
{
public:
	ListNode *Merge(ListNode *pHead1, ListNode *pHead2)
	{
		if (pHead1 == NULL)
			return pHead2;
		if (pHead2 == NULL)
			return pHead1;

		ListNode mergeNode(0);
		// 要记住保存头结点，最后返回
		ListNode *pNode = &mergeNode;

		while (pHead1 && pHead2)
		{
			if (pHead1->val < pHead2->val)
			{
				pNode->next = pHead1;
				pHead1 = pHead1->next;
			}
			else
			{
				pNode->next = pHead2;
				pHead2 = pHead2->next;
			}
			// 更新
			pNode = pNode->next;
		}
		if (pHead1)
			pNode->next = pHead1;
		if (pHead2)
			pNode->next = pHead2;

		return mergeNode.next;
	}

	ListNode *ReverseList(ListNode *pHead)
	{
		if (pHead == NULL)
			return NULL;

		ListNode *pReverse = NULL;
		ListNode *pNext = NULL;
		// 不是pHead->next
		while (pHead)
		{
			pNext = pHead->next;	// 保存next，防止原链表断开
			pHead->next = pReverse; // 修改pHead指向新链表
			pReverse = pHead;		// 更新新链表头指针，为下次做准备
			pHead = pNext;			// 旧链表移动至下一个节点
		}

		return pReverse;
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
};

int main()
{
	Solution solve;
	ListNode *list1 = NULL;
	ListNode *list2 = NULL;
	for (int i = 1; i < 10; i += 2)
		solve.addToTail(&list1, i);
	solve.printNode(list1);
	for (int i = 2; i < 8; i += 2)
		solve.addToTail(&list2, i);
	solve.printNode(list2);

	ListNode *head = solve.ReverseList(list1);
	while (head)
	{
		printf("%d ", head->val);
		head = head->next;
	}

	/*
	ListNode n1(1);
	ListNode n2(1);
	ListNode n3(1);
	ListNode n4(3);
	ListNode n5(3);
	ListNode n6(6);

	n1.next = &n2;
	n2.next = &n3;
	n3.next = &n4;
	n4.next = &n5;
	n5.next = &n6;

	solve.printNode(&n1);
	*/
	return 0;
}
