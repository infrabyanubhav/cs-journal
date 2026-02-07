#include<iostream>
using namespace std;



/**
@brief: assignment for dynamic memory allocation
 **/

class demo{
    public:
    int* num_1;
    int* num_2;

    demo(int one, int two){
        num_1 = new int;
        num_2 = new int;
        *num_1=one;
        *num_2=two;

    }

    int add(){
        cout<<"Output: "<<*num_1+*num_2<<endl;
        return 0;
    }
};



int main(){
    int a,b;
    cout<<"Enter value for a: "<<endl;
    cin>>a;
    cout<<"Enter value for b: "<<endl;
    cin>>b;
    demo *d1=new demo(a,b);
    d1->add();
    delete d1;
    return 0;


};