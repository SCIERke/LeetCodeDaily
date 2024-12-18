#include <iostream>
#include <vector>
#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    void fixPlusTen(vector<int>& digits) {
        for (int i=digits.size()-1 ;i >= 0 ;i--) {
            if (digits[i] >= 10) {
                digits[i] = digits[i] - 10;
                if (i-1 >= 0) {
                    digits[i-1] += 1; 
                } else {
                    digits.emplace(digits.begin(), 1); 
                }
            }
        }
    }       
    vector<int> plusOne(vector<int>& digits) {
        digits[digits.size() -1 ] +=1 ;
        fixPlusTen(digits);
        return digits;
    }
};