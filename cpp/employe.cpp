#include<iostream>
using namespace std;

class employee{
    public:
    int emp_uid;
    string emp_name;

    public:
    employee(int emp_uid, string emp_name){
        this->emp_uid=emp_uid;
        this->emp_name=emp_name;
        cout<<"emloyee_id"<<"\t"<<emp_uid<<endl;
        cout<<"employee_name"<<"\t"<<emp_name<<endl;

    };



};


int main(){
    employee e1(12345,"Noni");
    return 0;
}