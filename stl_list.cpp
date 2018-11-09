//测试stl中list常见API
//链表内存是非连续的，只有在需要时才分配内存

#include <iostream>
#include <list>
#include <string>
#include <algorithm>
using namespace std;

//打印
void printList(list<int> &mylist) {
	for(list<int>::iterator it = mylist.begin(); it != mylist.end();++it) {
		cout << *it << " ";
	}
	cout << endl;
}

//list容器初始化
void test01() {
	list<int> mylist; //默认构造
	list<int> mylist2(10,5); //带参数的构造函数
	list<int> mylist3(mylist2.begin(),mylist2.end());
	list<int> mylist4(mylist3);
}

//list容器插入删除
void test02() {
	list<int> mylist; 
	mylist.push_back(10);//尾插法
	mylist.push_back(20);
	mylist.push_front(30);//头插法
	mylist.push_front(40);
	mylist.insert(mylist.begin(),50);
	mylist.push_front(60);

	printList(mylist);

	//删除
	mylist.remove(40);
	printList(mylist);

	mylist.pop_back();
	mylist.pop_front();
	mylist.erase(mylist.begin());
	printList(mylist);

	mylist.erase(mylist.begin(),mylist.end());
	cout << "size:" << mylist.size() << endl;

}

//size大小和赋值
void test03() {
	list<int> mylist;
	cout << "size:" << mylist.size() << endl;
	for(int i=0;i < 10 ;++i) {
		mylist.push_back(i);
	}
	cout << "size:" << mylist.size() << endl;

	if(mylist.empty()) {
		cout << "empty" << endl;
	}else {
		cout << "not empty" << endl;
	}

	//assign
	list<int> mylist2;
	mylist2.assign(mylist.begin(),mylist.end());
	cout << "mylist:";
	printList(mylist);
	cout << "mylist2:";
	printList(mylist2);

	list<int> mylist3;
	mylist3 = mylist2; //operator=赋值
	cout << "mylist3:";
	printList(mylist3);

	//reverse
	mylist3.reverse();
	cout << "mylist3 reverse:";
	printList(mylist3);

	//swap
    mylist.swap(mylist3);
	cout << "after swap,mylist:";
	printList(mylist);

	
}

bool mycompare(int val1, int val2) {
	return val1 > val2;
}
//list排序
void test04() {
	list<int> mylist;
	mylist.push_back(1);
	mylist.push_back(8);
	mylist.push_back(5);
	mylist.push_back(9);

	printList(mylist);
	mylist.sort(); //默认从小到大
	printList(mylist);
	mylist.sort(mycompare);
	printList(mylist);
}

class Student {
public:
	Student(string _name, int _age):name(_name),age(_age){}
	string name;
	int age;
};

bool compareStu(Student obj1, Student obj2) {
	if(obj1.age > obj2.age) {
		return true;
	}	
	else {
		return false;
	}
}

void test05() {
	list<Student> mylist;
	Student s1("aaa",10),s2("bbb",2),s3("ccc",25);
	mylist.push_back(s1);
	mylist.push_back(s2);
	mylist.push_back(s3);

	for(list<Student>::iterator it=mylist.begin();it != mylist.end();++it) {
		cout << it->name << " " << it->age << endl;
	}

	mylist.sort(compareStu);

	for(list<Student>::iterator it=mylist.begin();it != mylist.end();++it) {
		cout << it->name << " " << it->age << endl;
	}


}
int main() {
	//test02();
	//test03();
	//test04();
	test05();
	return 0;
}
