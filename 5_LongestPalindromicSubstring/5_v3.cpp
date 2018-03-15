// Approach #4 (Expand Around Center)
class Solution {
public:
    string longestPalindrome(string s) {
        int maxLen = 0;
        int maxStart = 0;
        int slen = s.size();
        for (int i = 0; i < slen; i++) {
            int len1 = expandLen(s, slen, i, i);
            int len2 = expandLen(s, slen, i, i+1);
            
            int len = (len1 > len2)? len1 : len2;
            if (len > maxLen) {
                maxLen = len;
                maxStart = i - (len-1)/2;
            }
        }
        return s.substr(maxStart, maxLen);
    }
    
    int expandLen(const string& s, int slen, int l, int r) {
        while (l >= 0 && r < slen && s[l] == s[r]) {
            l--;
            r++;
        }
        return r - l - 1;
    }
};
