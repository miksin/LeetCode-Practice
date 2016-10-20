public class Solution {
    public int lengthOfLongestSubstring(String s) {
        int max_len = 0;
        HashMap<String, Integer> char_set = new HashMap<String, Integer>();

        for (int i = 0, start = 0; i < s.length(); i++){
            char c = s.charAt(i);
            
            Integer idx = char_set.get("" + c);
            if (idx != null){
                while (start < idx + 1){
                    char_set.remove("" + s.charAt(start));
                    start++;
                }
            }

            max_len = Math.max(i - start + 1, max_len);
            char_set.put("" + c, i);
        }

        return max_len;
    }
}