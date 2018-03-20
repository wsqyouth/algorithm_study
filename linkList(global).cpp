//这份代码仅仅起参考作用，由于使用的是一级指针，不是很规范
#include <stdio.h>
#include <stdlib.h>
#include <assert.h>
#include <stack>
using namespace std;

//s一般表示新生节点,p表示扫描指针

/* 结构体定义 */
struct ListNode
{
	int value;
	ListNode *next;
};

/* 链表头 */
ListNode* head = NULL;

/*尾插法插入值*/
void insertNode(int val)
{
	if (head == NULL)
	{
		ListNode *s = (ListNode*)malloc(sizeof(ListNode));
		s->value = val;
		s->next = NULL;

		head = s;
		return;
	}
	else
	{
		//查找尾结点
		ListNode *p = head;
		while (p->next != NULL)
		{
			p = p->next;
		}

		ListNode *s = (ListNode*)malloc(sizeof(ListNode));
		s->value = val;
		s->next = NULL;
		//插入新生节点
		s->next = p->next;
		p->next = s;
	}
}
/* 删除第一个value=val的链表节点 */
void deleteNode(int val)
{
	ListNode *curr = head;
	ListNode *prev = NULL;

	while (curr != NULL)
	{
		if (curr->value == val)
		{
			/* 如果待删除的节点是链表头 */
			if (curr == head)
			{
				head = head->next;
			}
			else
			{
				prev->next = curr->next;
			}
			free(curr);
			return;
		}
		//存储previous指针并遍历所有节点
		prev = curr;
		curr = curr->next;
	}
}
/*根据头结点打印(不需要修改头结点地址，因此一级指针)*/
void printNode(ListNode *myhead)
{
	if (myhead == NULL) return;
	while (myhead)
	{
		printf("%d ", myhead->value);
		myhead = myhead->next;
	}
	printf("\n");
}
/* 删除value=val的链表节点 */
int delete_from_list(int val)
{
	ListNode ** curr = &head; //正在遍历的节点的指针  
	ListNode * entry; //正在遍历的节点  

	while (*curr)
	{
		entry = *curr;
		if (entry->value == val)
		{
			/* 删除entry节点 */
			*curr = entry->next;
			free(entry);
			return 0;
		}
		/* 遍历所有节点的指针 */
		curr = &(entry->next);

	}
	return -1;
}

void printReverseRecur(ListNode *pHead)
{
	if (pHead != NULL)
	{
		if (pHead->next != NULL)
		{
			printReverseRecur(pHead->next);
		}
		printf("%d\t", pHead->value);
	}
}

void printListReversingly(ListNode *pHead)
{
	stack<ListNode *> nodeStack;
	ListNode *pNode = pHead;
	while (pNode != NULL)
	{
		nodeStack.push(pNode);
		pNode = pNode->next;
	}
	while (!nodeStack.empty())
	{
		pNode = nodeStack.top();
		printf("%d\t", pNode->value);
		nodeStack.pop();
	}
}

ListNode* FindKthToTail(ListNode* pListHead, unsigned int k)
{
	if (pListHead == NULL) return NULL;
	if (k < 1) return NULL;
	//若k大于结点数，返回NULL
	ListNode * pNode = pListHead;
	unsigned int count = 0;
	while (pNode != NULL)
	{
		pNode = pNode->next;
		count++;
	}
	if (k > count) return NULL;
	//快慢指针都指向head
	unsigned int i = 0;
	ListNode * pNodeFast = pListHead;
	ListNode * pNodeSlow = pListHead;

	while (pNodeFast != NULL && i < k)
	{
		pNodeFast = pNodeFast->next;
		i++;
	}
	//fast走k步后，slow开始走
	while (pNodeFast != NULL)
	{
		pNodeFast = pNodeFast->next;
		pNodeSlow = pNodeSlow->next;
	}

	return pNodeSlow;
}
int main()
{
	//插入
	for (int i = 1; i < 9; ++i)
	{
		insertNode(i);
	}
	//打印
	printNode(head);
	if (FindKthToTail(head, 4) !=NULL)
		printf("%d", FindKthToTail(head, 4)->value);
	
/*
	//反转打印
	printReverseRecur(head);
	printf("\n");
	printListReversingly(head);
	
	//删除(一级指针)
	deleteNode(2);
	printNode(head);
	//删除(二级指针)
	delete_from_list(1);
	printNode(head);
	*/
	return 0;
}
