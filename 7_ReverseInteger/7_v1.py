class Solution(object):
    def reverse(self, x):
        """
        :type x: int
        :rtype: int
        """
        stringValue = str(abs(x))
        reverseStr = stringValue[::-1]
        
        if x > 0 and int(reverseStr) <= 2147483647:
            return int(reverseStr)
        elif x < 0 and -int(reverseStr) >= -2147483648:
            return -int(reverseStr)
        
        return 0
