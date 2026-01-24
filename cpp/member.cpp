//College Assignment 
#include <iostream>
using namespace std;

class BankAccount {
private:
    double balance;   

public:
    int accountNumber;  

    void deposit(double amount) {
        balance = amount;
    }

    void showBalance() {
        cout << "Account Number: " << accountNumber << endl;
        cout << "Balance: ₹" << balance << endl;
    }
};

int main() {
    BankAccount acc;
    int accountNumber;
    cin>>accountNumber;
    acc.accountNumber = accountNumber;  
    acc.deposit(5000);          
    acc.showBalance();



    return 0;
}
