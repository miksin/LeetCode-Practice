class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        map = {}
        
        for i, x in enumerate(nums):
            if target - x in map:
                return [map[target - x], i]
            map[x] = i