#include <bits/stdc++.h>
#include <cmath>
using namespace std;

int findIndexofGivenChar(vector<char> &vec ,char target) {
	return find(vec.begin() , vec.end() ,target) - vec.begin();
}

int main() {
	string s;
	cin >> s;
	
	int value = 0;
	int len = s.length() - 1;
	int index = 0;

	vector<char> vec = { 'I' ,'V' ,'X' ,'L' ,'C' ,'D' ,'M'};
	map<char ,int> romanValue = {
		{'I', 1 },
		{'V', 5 },
		{'X', 10 },
		{'L', 50 },
		{'C', 100},
		{'D', 500}, 
		{'M', 1000},

	};
	
	while (index <= len) {
		int charIndex = findIndexofGivenChar(vec ,s[index]);
		if ((charIndex == 0 |charIndex == 2 | charIndex == 4) && index+1 <= len) {
			int diffOfCharIndex = abs(charIndex  - findIndexofGivenChar(vec ,s[index+1]));
			if ( diffOfCharIndex == 1 || diffOfCharIndex == 2) {
				value += romanValue[s[index+1]] - romanValue[s[index]];
				index += 2;
			} else {
				value += romanValue[s[index]];
				index++;
			}
		} else {
			value += romanValue[s[index]];
			index++;
		}
	}
	cout << value;
}
