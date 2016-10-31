class Solution {
public:
    string convert(string s, int numRows) {
        int unitLen = (numRows==1)? 1: numRows*2 - 2;
        vector<string> rows(numRows);
        string result;
        
        for (int i = 0; i < s.size(); i++){
            int unitIdx = i % unitLen;
            int rowIdx = (unitIdx < numRows)? unitIdx: unitLen - unitIdx;
            rows[rowIdx] += s[i];
        }
        
        for (string& row: rows){
            result += row;
        }
        
        return result;
    }
};