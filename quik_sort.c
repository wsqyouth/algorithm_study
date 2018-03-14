//快速排序是不稳定算法，这里测试的快排看起来非常简单，经常复习之


#include <stdio.h> 
//确定一趟排序元素的位置
int identityTheIndex(int A[],int low,int high){
         
    int key = A[low];
    while(low<high){
        
        while(low<high&&A[high]>=key){
             high--;
        }
        A[low] = A[high];
        while(low<high&&A[low]<=key){
             low++;
        }
        A[high] = A[low];
    }
    A[low]=key;
        
    return low;
         
}
 
 
//递归进行排序
void quictSort(int A[],int low, int high){
    if(low < high){
         int index =  identityTheIndex(A,low,high);
         quictSort(A,low,index-1);
         quictSort(A,index+1,high);
     }
}
int main()
{   

    int arr[] = {3,2,9,8,2,1,6};
    quictSort(arr,0,6);
    for(int i=0;i<7;i++)
        printf("%d ",arr[i]);
    return 0;
}
