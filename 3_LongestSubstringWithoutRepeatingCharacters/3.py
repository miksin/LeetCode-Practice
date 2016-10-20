class Solution(object):
    def lengthOfLongestSubstring(self, s):
        """
        :type s: str
        :rtype: int
        """
        
        max_len = 0
        start = 0
        char_set = {}

        for i, c in enumerate(s):
            if c in char_set:
                new_start = char_set[c] + 1
                while start < new_start:
                    char_set.pop(s[start])
                    start += 1
                
            max_len = max(i - start + 1, max_len)
            char_set[c] = i

        return max_len
