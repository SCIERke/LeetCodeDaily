#include <iostream>
#include <vector>
#include <algorithm>
#include <cmath>

using namespace std;

class Solution {
public:
    long long findMostPilesIndex(vector<int>gifts_temp) {
        long long index = 0 , max_value = -1;

        for (long long i = 0 ; i < gifts_temp.size();i++) {
            if (gifts_temp[i] > max_value) {
            max_value = gifts_temp[i];
            index = i;
            }
        }
        return index;
    }
    long long pickGifts(vector<int>& gifts, int k) {
        long long sum = 0;

        for (int i = 0 ;i < k ;i++) {
            gifts[findMostPilesIndex(gifts)] = sqrt(gifts[findMostPilesIndex(gifts)]);
        }
        for (int i = 0 ; i < gifts.size() ;i++) {
            sum += gifts[i];
        }
        return sum;
    }
};