#include <iostream>
using namespace std;

// Copy Constructor
class CopyConstructor {
public:
    CopyConstructor() {
        cout << "Default Constructor Called" << endl;
    }

    CopyConstructor(const CopyConstructor &obj) {
        cout << "Copy Constructor Called" << endl;
    }
};

// Default Constructor
class DefaultConstructor {
public:
    DefaultConstructor() {
        cout << "Default Constructor Example" << endl;
    }
};

// Dynamic Constructor
class DynamicConstructor {
private:
    int *p;

public:
    DynamicConstructor(int value) {
        p = new int(value);
        cout << "Dynamic Constructor Called: " << *p << endl;
    }

    ~DynamicConstructor() {
        delete p;
    }
};

// Multiple Constructors
class MultipleConstructor {
public:
    MultipleConstructor() {
        cout << "Constructor with No Parameters" << endl;
    }

    MultipleConstructor(int a) {
        cout << "Constructor with One Parameter: " << a << endl;
    }

    MultipleConstructor(int a, int b) {
        cout << "Constructor with Two Parameters: "
             << a << ", " << b << endl;
    }
};

// Parameterized Constructor with Default Arguments
class ParamConstructor {
public:
    ParamConstructor(int a = 10, int b = 20) {
        cout << "Parameterized Constructor Called: "
             << a << ", " << b << endl;
    }
};

int main() {
    cout << "Demonstration of Different Types of Constructors\n\n";

    // Copy Constructor
    CopyConstructor c1;
    CopyConstructor c2 = c1;

    cout << endl;

    // Default Constructor
    DefaultConstructor d1;

    cout << endl;

    // Dynamic Constructor
    DynamicConstructor dc(100);

    cout << endl;

    // Multiple Constructors
    MultipleConstructor m1;
    MultipleConstructor m2(10);
    MultipleConstructor m3(10, 20);

    cout << endl;

    // Parameterized Constructor
    ParamConstructor p1;

    return 0;
}
