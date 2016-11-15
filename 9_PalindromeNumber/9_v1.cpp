class Solution {
public:
    bool isPalindrome(int x) {
        int base = 10;
        int originInt = x;
        int reverseInt = 0;
        
        if (x < 0)
            return false;
        
        while (x > 0){
            int digit = x % base;
            reverseInt = reverseInt*base + digit;
            x /= base;
        }
        
        return originInt == reverseInt;
    }
};
