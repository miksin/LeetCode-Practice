class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        int max_len = 0;
        int char_set[256];
        for (int& x : char_set){
            x = -1;
        }


        for (int i = 0, start = 0; i < s.length(); i++){
            char c = s[i];
            
            if (char_set[c] != -1){
                while(start < char_set[c] + 1){
                    char_set[s[start++]] = -1;
                }
            }
            max_len = max(i - start + 1, max_len);
            char_set[c] = i;
        }

        return max_len;
    }
};