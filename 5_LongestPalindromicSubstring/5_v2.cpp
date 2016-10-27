#include <string>
using namespace std;

class Solution {
public:
    string longestPalindrome(string s) {
        int MAX_SIZE = 1000;
        int len = s.size();
        bool table[MAX_SIZE][MAX_SIZE];

        int max_sublen = 0;
        string max_substr;

        for (int i = len - 1; i >= 0; i--){
            for (int j = 0; j < len; j++){
                table[i][j] = (j <= i) || (table[i + 1][j - 1] && s[i] == s[j]);
                if (table[i][j] && j - i + 1 > max_sublen){
                    max_sublen = j - i + 1;
                    max_substr = string(s, i, max_sublen);
                }
            }
        }

        return max_substr;
    }
};