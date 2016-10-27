// Time Limit Exceeded

public class Solution {
    public String longestPalindrome(String s) {

        int len = s.length();
        boolean[][] table = new boolean[len][len];

        int max_len = 0;
        String max_str = new String();

        for (int i = len - 1; i >= 0; i--){
            for (int j = 0; j < len; j++){
                table[i][j] = (j <= i) || (table[i + 1][j - 1] && s.charAt(i) == s.charAt(j));

                if (table[i][j] && j - i + 1 > max_len){
                    max_len = j - i + 1;
                    max_str = s.substring(i, j + 1);
                }
            }
        }

        return max_str;
    }
}