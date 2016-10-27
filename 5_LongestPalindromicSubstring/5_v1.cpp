#include <string>
#include <vector>
using namespace std;

// Time Limit Exceeded

class Solution {
public:
    string longestPalindrome(string s) {
        int len = s.size();
        vector<vector<bool> > table(len, vector<bool>(len, true));

        int max_sublen = 0;
        string max_substr;

        for (int i = len - 1; i >= 0; i--){
            for (int j = i; j < len; j++){
                table[i][j] = (i == j) || (table[i + 1][j - 1] && s[i] == s[j]);
                if (table[i][j] && j - i + 1 > max_sublen){
                    max_sublen = j - i + 1;
                    max_substr = string(s, i, max_sublen);
                }
            }
        }

        return max_substr;
    }
};