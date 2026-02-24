#include<iostream>
using namespace std;

class Sum{
    float num_1,num_2;

    public:
    Sum(float n1, float n2){

        num_1=n1;
        num_2=n2;

    };

    friend void CalculateSum(Sum obj);
};



void CalculateSum(Sum obj){
    cout<<"The Sum is = "<<obj.num_1+obj.num_2;
};



int main(){
    float a,b;
    cout<<"Enter the First Number "<<endl;
    cin>>a;
    cout<<"Enter the Second Number "<<endl;
    cin>>b;
    Sum obj(a,b);
    CalculateSum(obj);

    return 0;
}