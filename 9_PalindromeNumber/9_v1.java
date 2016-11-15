public class Solution {
    public boolean isPalindrome(int x) {
        int base = 10;
        int originInt = x;
        int reverseInt = 0;
        
        if (x < 0)
            return false;
            
        while (x > 0){
            reverseInt = reverseInt*base + x%base;
            x /= base;
        }
        
        return reverseInt == originInt;
    }
}
