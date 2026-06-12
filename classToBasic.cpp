#include <iostream>
using namespace std;

class Time{
   
    int hrs;
    int min;
    public:
    Time(int a, int b){
        hrs=a;
        min=b;
    };
     operator int(){
        cout<<"Conversion of basic to class type t"<<endl;
        return (hrs*60+min);
    }
};

int main(){
    cout<<"Executing....."<<endl;
    int h,m, duration;
    h=4;
    m=40;
    Time t1(h,m);
    duration= t1;
    cout<<duration<<endl;
    duration=t1.operator int();
    cout<<duration;
    return 0;
}
