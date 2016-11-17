class Solution(object):
    def longestCommonPrefix(self, strs):
        """
        :type strs: List[str]
        :rtype: str
        """
        if len(strs) == 0:
            return ""
        
        for i, c in enumerate(strs[0]):
            for str in strs:
                if len(str) <= i or str[i] != c:
                    return strs[0][0:i]
        
        return strs[0]
            
