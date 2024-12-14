#include <stack>
#include <iostream>
using namespace std;

class Solution {
public:
    bool isValid(string s) {
        stack<char> prn;
        for (int i = 0; i < s.size(); i++) {
            if (s[i] == '(' || s[i] == '[' || s[i] == '{') {
                prn.push(s[i]);
            } else {
                if (prn.empty()) {
                    return false;
                }
                if (s[i] == ')' && prn.top() == '(') {
                    prn.pop();
                } else if (s[i] == '}' && prn.top() == '{') {
                    prn.pop();
                } else if (s[i] == ']' && prn.top() == '[') {
                    prn.pop();
                } else {
                    return false;
                }
            }
        }

        return prn.empty();
    }
};