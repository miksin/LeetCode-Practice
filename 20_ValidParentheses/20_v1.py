class Solution(object):
    def isValid(self, s):
        """
        :type s: str
        :rtype: bool
        """
        mapping = {
            ')': '(',
            ']': '[',
            '}': '{',
        }
        
        check = []
        
        for c in s:
            if c in mapping:
                if not check or mapping[c] != check[-1]:
                    return False
                check.pop()
            else:
                check.append(c)
        
        return not check
