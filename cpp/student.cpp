#include<iostream>
using namespace std;

class student_name{
    public:
    int roll_no;
    int marks;
    void student_info();
};

void student_name::student_info(){
    cout<<"student roll_no. "<<roll_no<<"\n";
    cout<<"student marks"<<marks;
};


int main(){

    student_name first;
    first.roll_no=2905;
    first.marks=95;

    first.student_info();

    return 0;

    
}






