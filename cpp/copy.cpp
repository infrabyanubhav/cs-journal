#include <iostream>
using namespace std;

/** 

@brief: enter and print the details of student
@brief:using copy constructor mechanism.
@param:string:name, string:roll_no.

 **/

class Student{
    public:
    string name;
    string roll_no;

    Student(string s_name,string s_roll_no){
        name=s_name;
        roll_no=s_roll_no;
    }

    Student(const Student &s){
        cout<<"student name: "<<s.name<<endl;
        cout<<"student roll_no: "<<s.roll_no<<endl;
    }

};


int main(){
    string s_name,s_roll_number;
    cout<<"Enter student name"<<endl;
    cin>>s_name;
    cout<<"Enter student roll_no."<<endl;
    cin>>s_roll_number;

    Student student_1(s_name,s_roll_number);
    Student student_2=student_1;

    return 0;
    
}