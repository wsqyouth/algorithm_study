#include <iostream>
#include <stack>
#include <map>
using namespace std;
class Solution
{
public:
  bool isValid(string s)
  {
    if (s.empty())
    {
      return false;
    }
    stack<int> st;
    map<char, char> hash = {{'}', '{'}, {')', '('}, {']', '['}};

    for (int i = 0; i < s.length(); i++)
    {
      if (s[i] == '(' || s[i] == '[' || s[i] == '{')
      {
        st.push(s[i]);
      }
      else if (s[i] == ')' || s[i] == '}' || s[i] == ']')
      {
        if (st.empty() || st.top() != hash[s[i]])
        {
          return false;
        }
        else
        {
          st.pop();
        }
      }
    }
    return st.empty();
  }
};

int main()
{
  string str = "()[]}";
  Solution objs;
  cout << objs.isValid(str) << endl;
  return 0;
}

/*
总结：
g++ 20_Valid_Parentheses.cpp -o 20_Valid_Parentheses -std=c++11 && ./20_Valid_Parentheses
1. 熟悉cpp栈的用法 http://www.cplusplus.com/reference/stack/stack/?kw=stack
2. 犯错：未考虑pop时必须栈顶为对应元素、未考虑pop时st为空则直接返回false、结合map使用时另一半对应问题
*/
