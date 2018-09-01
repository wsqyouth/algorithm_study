
调整数组顺序使奇数位于偶数前面

输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有的奇数位于数组的前半部分，所有的偶数位于数组的后半部分，
并保证奇数和奇数，偶数和偶数之间的相对位置不变。

参考牛客网的思路记录下：
链接：https://www.nowcoder.com/questionTerminal/beb5aa231adc45b2a5dcc5b62c93f593
来源：牛客网

//两个思路吧，第一个思路：类似冒泡算法，前偶后奇数就交换：
class Solution {
public:
    void reOrderArray(vector<int> &array) {
 
         
        for (int i = 0; i < array.size();i++)
        {
            for (int j = array.size() - 1; j>i;j--)
            {
                if (array[j] % 2 == 1 && array[j - 1]%2 == 0) //前偶后奇交换
                {
                    swap(array[j], array[j-1]);
                }
            }
        }
    }
};
 
//第二个思路：再创建一个数组
class Solution{
public:
    void reOrderArray(vector<int> &array) {
 
        vector<int> array_temp;
        vector<int>::iterator ib1, ie1;
        ib1 = array.begin();
 
 
        for (; ib1 != array.end();){            //遇见偶数，就保存到新数组，同时从原数组中删除
            if (*ib1 % 2 == 0) {
                array_temp.push_back(*ib1);
                ib1 = array.erase(ib1);
            }
            else{
                ib1++;
            }
 
        }
        vector<int>::iterator ib2, ie2;
        ib2 = array_temp.begin();
        ie2 = array_temp.end();
 
        for (; ib2 != ie2; ib2++)             //将新数组的数添加到老数组
        {
            array.push_back(*ib2);
        }
    }
};


//第三个思路：
用的STL　stable_partition 这个函数
函数功能是将数组中　isOk为真的放在数组前，假的放在数组后，和题意相符
链接：https://www.nowcoder.com/questionTerminal/beb5aa231adc45b2a5dcc5b62c93f593
来源：牛客网
bool isOk(int n){  return ((n & 1) == 1);　//奇数返回真 }
class Solution{
    void reOrderArray(vector<int> &array){
        stable_partition(array.begin(),array.end(),isOk);
    }
};
