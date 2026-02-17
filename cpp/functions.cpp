#include<iostream>

using namespace std;


/**
@brief: Write a programme using functions sum, multiply, divide and subtract two values;
@param:(float a, float b)
*/


//Sum
float sum(float a,float b){ 
    cout<<"Sum"<<endl;
    return a+b;
}

//multiply
float multiply(float a,float b){
    cout<<"multiply"<<endl;
    return a*b;
}

//divide
float divide(float a,float b){
    cout<<"divide"<<endl;
    return a/b;
}

//sub
float sub(float a,float b){
    cout<<"sub"<<endl;
    return a-b;
}


int main(){
float a,b;
float retunedValue;     //returnedValue var for capturing return
cout<<"Enter first Number"<<endl;
cin>>a;
cout<<"enter second number"<<endl;
cin>>b;

retunedValue=sum(a,b);
cout<<retunedValue<<endl;

retunedValue=divide(a,b);
cout<<retunedValue<<endl;

retunedValue=sub(a,b);
cout<<retunedValue<<endl;

retunedValue=multiply(a,b);
cout<<retunedValue<<endl;


}



