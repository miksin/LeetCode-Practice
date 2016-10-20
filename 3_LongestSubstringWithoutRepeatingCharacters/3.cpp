// Time Limit Exceeded

class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        int max_len = 0;
        map<char, int> char_set;

        for (int i = 0, start = 0; i < s.length(); i++){
            char c = s[i];
            map<char, int>::iterator it(char_set.find(c));

            if (it != char_set.end()){
                max_len = max(i - start, max_len);

                while(start < it->second + 1){
                    char_set.erase(char_set.find(s[start++]));
                }
            }
            char_set[c] = i;
        }

        return max_len;
    }
};