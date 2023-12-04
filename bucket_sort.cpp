// 桶排序代码
http : // www.cnblogs.com/hdk1993/p/4671806.html

#include <iostream>
#include <cstdlib>
#include <cstdio>
       using namespace std;

void bucketSort(double *a, int n)
{
    // 链表结点描述
    typedef struct Node
    {
        double key;
        struct Node *next;
    } Node;
    // 辅助数组元素描述
    typedef struct
    {
        Node *next;
    } Head;
    int i, j;
    Head head[10] = {NULL};
    Node *p;
    Node *q;
    Node *node;
    for (i = 1; i <= n; i++)
    {
        node = (Node *)malloc(sizeof(Node));
        node->key = a[i];
        node->next = NULL;
        p = q = head[(int)(a[i] * 10)].next;
        if (p == NULL)
        {
            head[(int)(a[i] * 10)].next = node;
            continue;
        }
        while (p)
        {
            if (node->key < p->key)
                break;
            q = p;
            p = p->next;
        }
        if (p == NULL)
        {
            q->next = node;
        }
        else
        {
            node->next = p;
            q->next = node;
        }
    }
    j = 1;
    for (i = 0; i < 10; i++)
    {
        p = head[i].next;
        while (p)
        {
            a[j++] = p->key;
            p = p->next;
        }
    }
}

int main()
{
    int i;
    double a[13] = {0, 0.13, 0.25, 0.18, 0.29, 0.81, 0.52, 0.52, 0.83, 0.52, 0.69, 0.13, 0.16}; // 不考虑a[0]
    bucketSort(a, 12);
    for (i = 1; i <= 12; i++)
        printf("%-6.2f", a[i]);
    printf("\n");
    return 0;
}
