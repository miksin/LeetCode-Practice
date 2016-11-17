class Solution(object):
    def romanToInt(self, s):
        """
        :type s: str
        :rtype: int
        """
        alpha = {
            'I': 1,
            'V': 5,
            'X': 10,
            'L': 50,
            'C': 100,
            'D': 500,
            'M': 1000,
        }
        
        buf = 0
        sum = 0
        
        for c in s:
            if buf == 0:
                buf = alpha[c]
            elif buf < alpha[c]:
                sum = sum - buf + alpha[c]
                buf = 0
            else:
                sum = sum + buf
                buf = alpha[c]
            
        sum = sum + buf
        return sum
