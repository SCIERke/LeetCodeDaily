#include <vector>
#include <queue>
#include <climits>
using namespace std;

class Solution {
public:
    long long findScore(vector<int>& nums) {
        long long score = 0;
        int n = nums.size();
        vector<bool> marked(n, false);

        priority_queue<pair<int, int>, vector<pair<int, int>>, greater<>> pq;
        for (int i = 0; i < n; i++) {
            pq.emplace(nums[i], i);
        }

        while (!pq.empty()) {
            auto [value, index] = pq.top();
            pq.pop();

            if (marked[index]) continue;

            score += value;
            marked[index] = true;
            if (index - 1 >= 0) marked[index - 1] = true;
            if (index + 1 < n) marked[index + 1] = true;
        }

        return score;
    }
};
