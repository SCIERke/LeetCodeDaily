#include <iostream>

using namespace std;

int cnt = 0;

void sol(int n,int score) {
  if (score == n) {
    cnt++;
    return;
  } else if (score > n) {
    return;
  } else {
    sol(n ,score + 1);
    sol(n , score + 2);
  }
}

int main() {
  int n;
  cin >> n;

  int start_score = 0;
  sol(n ,start_score);
  cout << cnt;
}