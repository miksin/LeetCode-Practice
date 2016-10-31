public class Solution {
    public String convert(String s, int numRows) {
        int unitLen = (numRows == 1)? 1: numRows*2 - 2;
        String[] rows = new String[numRows];
        String result = new String();
        
        for (int i = 0; i < numRows; i++){
            rows[i] = new String();
        }
        
        for (int i = 0; i < s.length(); i++){
            int unitIdx = i % unitLen;
            int rowIdx = (unitIdx < numRows)? unitIdx: unitLen - unitIdx;
            rows[rowIdx] += s.charAt(i);
        }
        
        for (int i = 0; i < numRows; i++){
            result += rows[i];
        }
        
        return result;
    }
}