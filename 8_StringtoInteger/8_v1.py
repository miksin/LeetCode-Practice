class Solution(object):
    def myAtoi(self, str):
        """
        :type str: str
        :rtype: int
        """
        INT_MAX = 2147483647
        INT_MIN = -2147483648
        
        m = re.match("^ *([\\+-]?[0-9]+).*$", str)
        
        if not m:
            return 0
        
        result = int(m.group(1))
        
        return min(result, INT_MAX) if result > 0 else max(result, INT_MIN)
