class Solution(object):
    def convert(self, s, numRows):
        """
        :type s: str
        :type numRows: int
        :rtype: str
        """
        
        unitLen = 1 if numRows == 1 else numRows*2 - 2
        rows = [''] * numRows
        
        for i, c in enumerate(s):
            unitIdx = i % unitLen
            rowIdx = unitIdx if unitIdx < numRows else unitLen - unitIdx
            rows[rowIdx] = rows[rowIdx] + c
        
        return "".join(rows)
        