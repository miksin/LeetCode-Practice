class Solution {
public:
    string longestCommonPrefix(vector<string>& strs) {
        if (strs.size() == 0)
            return "";
        
        for (int i = 0; strs[0].size(); i++){
            char c = strs[0][i];
            for (int j = 1; j < strs.size(); j++){
                if (strs[j].size() <= i || c != strs[j][i])
                    return string(strs[0].begin(), strs[0].begin() + i);
            }
        }
        
        return strs[0];
    }
};
