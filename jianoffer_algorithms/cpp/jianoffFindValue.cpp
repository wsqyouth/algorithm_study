#include <stdio.h>
#include <stdlib.h>
#include <assert.h>

/* 结构体定义 */
struct ListNode
{
	int value;
	ListNode *next;
};

/*根据头结点打印(不需要修改头结点地址，因此一级指针)*/
void printNode(ListNode *myhead)
{
	if (myhead == NULL)
		return;
	while (myhead)
	{
		printf("%d ", myhead->value);
		myhead = myhead->next;
	}
	printf("\n");
}

ListNode *FindKthToTail(ListNode *pListHead, unsigned int k)
{
	if (pListHead == NULL)
		return NULL;
	if (k < 1)
		return NULL;
	// 若k大于结点数，返回NULL
	ListNode *pNode = pListHead;
	unsigned int count = 0;
	while (pNode != NULL)
	{
		pNode = pNode->next;
		count++;
	}
	if (k > count)
		return NULL;
	// 快慢指针都指向head
	unsigned int i = 0;
	ListNode *pNodeFast = pListHead;
	ListNode *pNodeSlow = pListHead;

	while (pNodeFast != NULL && i < k)
	{
		pNodeFast = pNodeFast->next;
		i++;
	}
	if (pNodeFast != NULL)
	{
		// fast走k步后，slow开始走
		while (pNodeFast != NULL)
		{
			pNodeFast = pNodeFast->next;
			pNodeSlow = pNodeSlow->next;
		}

		return pNodeSlow;
	}

	return NULL;
}
void addtoTail(ListNode **pHead, int val)
{
	ListNode *pNew = (ListNode *)malloc(sizeof(ListNode));
	pNew->value = val;
	pNew->next = NULL;

	if (*pHead == NULL)
	{
		*pHead = pNew;
	}
	else
	{
		ListNode *pNode = *pHead;
		// 找尾结点
		while (pNode->next != NULL)
		{
			pNode = pNode->next;
		}
		pNew->next = pNode->next;
		pNode->next = pNew;
	}
}
// 去除值(Linus推荐)
void removeNode(ListNode **pHead, int val)
{

	if (pHead == NULL || *pHead == NULL)
	{
		return;
	}
	ListNode **curr = pHead; // 正在遍历的节点的指针
	ListNode *entry;		 // 正在遍历的节点
	while (*curr)
	{
		entry = *curr; // 注2
		if (entry->value == val)
		{
			/* 删除entry节点 */
			*curr = entry->next; // 注3
			free(entry);
			return;
		}
		/* 遍历所有节点的指针 */
		curr = &(entry->next); // 注1
	}
}
// 去除值（剑指offer推荐）
void RemoveNode(ListNode **pHead, int val)
{

	if (pHead == NULL || *pHead == NULL)
	{
		return;
	}
	ListNode *pToBeDeleted = NULL;
	if ((*pHead)->value == val)
	{
		pToBeDeleted = *pHead;
		*pHead = (*pHead)->next;
	}
	else
	{
		ListNode *pNode = *pHead;
		while (pNode->next != NULL && pNode->next->value != val)
			pNode = pNode->next;

		if (pNode->next != NULL && pNode->next->value == val)
		{
			pToBeDeleted = pNode->next;
			pNode->next = pToBeDeleted->next;
		}
	}

	if (pToBeDeleted != NULL)
	{
		free(pToBeDeleted);
	}
}
int main()
{
	ListNode *pHead = NULL;
	for (int i = 1; i < 9; i++)
		addtoTail(&pHead, i);
	printf("add to tail:\n");
	printNode(pHead);
	printf("find the last 4th value:%d\n", FindKthToTail(pHead, 4)->value);
	printf("remove value==1:\n");
	RemoveNode(&pHead, 1);
	printNode(pHead);
	return 0;
}
