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
    int getHrs() const{
        return hrs;
    };
    int getMin() const{
        return min;
    };
    };

class Duration{
    int duration;
    public:
    Duration(){};
     Duration(const Time& t){
        cout<<"Conversion of class to class type t"<<endl;
        duration= (t.getHrs() * 3600) + (t.getMin() * 60);
    }
    void display() { cout << "Total Seconds: " << duration << endl; }
};

int main(){
    cout<<"Executing....."<<endl;
    int h,m;
    Duration duration;
    h=4;
    m=40;
    Time t1(h,m);
    duration = t1;
    duration.display();
   
    return 0;
}
