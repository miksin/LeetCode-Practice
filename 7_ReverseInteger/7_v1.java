public class Solution {
    public int reverse(int x) {
        StringBuilder stringValue = new StringBuilder(Integer.toString(Math.abs(x)));
        String reverseStr = stringValue.reverse().toString();
        int result = 0;
        
        try {
            result = Integer.parseInt(reverseStr);
        } catch(Exception e) {
        }
        
        return result * (int)Math.signum(x);
    }
}
