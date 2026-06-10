#include <iostream>
using namespace std;

void callbyref(int *a, int *b) {
    (*a)++;
    (*b)++;
}

void callbyval(int e, int d) {
    e++;
    d++;
}

int main() {
    int a = 10, b = 20, e = 10, d = 20;

    callbyref(&a, &b);
    callbyval(e, d);

    cout << a << " " << b << " " << e << " " << d;
}
