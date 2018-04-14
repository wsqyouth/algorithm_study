//本代码使用递归的方法依次实现了前序遍历、中序遍历、后序遍历、层次遍历
//
树的形状：   8
        3         10
    1     6  null     15

各种遍历结果：
前序遍历：8 3 1 6 10 15 
中序遍历：1 3 6 8 10 15 
后序遍历：1 6 3 15 10 8 
层次遍历：8 3 10 1 6 15 

----

#include <iostream>
#include <vector>
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
   if(pHead == NULL) return;
   cout<<pHead->val<<" ";
   preOrder(pHead->left);
   preOrder(pHead->right);
}

void inOrder(TreeNode* pHead){
    //中序遍历，左 - 根 - 右
    if(pHead == NULL) return;
    inOrder(pHead->left);
    cout<<pHead->val<<" ";//当前节点在其左右子节点之间访问， 左 - 根 - 右
    inOrder(pHead->right);
}

void posOrder(TreeNode* pHead){
    //后序遍历，左 - 右 - 根
    if(pHead == NULL) return;
    posOrder(pHead->left);
    posOrder(pHead->right);
    cout<<pHead->val<<" ";//当前节点在其左右子节点之后访问，左-右-根
}
//从上往下打印出二叉树的每个节点，同层节点从左至右打印。
void levelOrder(TreeNode *pHead) {
   
    if(pHead==NULL) return ;
    
    deque<TreeNode *> dequeTreeNode;
    dequeTreeNode.push_back(pHead);
    while(!dequeTreeNode.empty())
    {
        TreeNode *temp = dequeTreeNode.front();
        //push val to vector
        cout<<temp->val<<" ";
        if(temp->left != NULL)
            dequeTreeNode.push_back(temp->left);
        if(temp->right != NULL)
            dequeTreeNode.push_back(temp->right);
        dequeTreeNode.pop_front();
    }
        
 }    
 
int main(){
    TreeNode a(8);
    TreeNode b(3);
    TreeNode c(10);
    TreeNode d(1);
    TreeNode e(6);
    TreeNode f(15);
    
    a.left = &b;
    a.right = &c;
    b.left = &d;
    b.right = &e;
    c.right = &f;
    
    preOrder(&a);
    cout<<endl;
    inOrder(&a);
    cout<<endl;
    posOrder(&a);
    cout<<endl;
    levelOrder(&a);
    cout<<endl;
	return 0;
}
