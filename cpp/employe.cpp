#include<iostream>
using namespace std;

class employee{
    public:
    int emp_uid;
    string emp_name;

    public:
    employee(){
        
        cout<<"emloyee_id"<<"\t"<<emp_uid<<endl;
        cout<<"employee_name"<<"\t"<<emp_name<<endl;

    };



};


int main(){
    employee e1;
    e1.emp_uid=12345;
    e1.emp_name="Noni";
    return 0;
}