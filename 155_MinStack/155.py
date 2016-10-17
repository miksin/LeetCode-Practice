class MinStack(object):
    
    def __init__(self):
        """
        initialize your data structure here.
        """
        self.stack = []
        self.minVal = sys.maxint

    def push(self, x):
        """
        :type x: int
        :rtype: void
        """
        if x < self.minVal:
            self.minVal = x
        self.stack.append(x)

    def pop(self):
        """
        :rtype: void
        """
        self.stack.pop()
        if len(self.stack) is not 0:
            self.minVal = min(self.stack)
        else:
            self.minVal = sys.maxint
        

    def top(self):
        """
        :rtype: int
        """
        return self.stack[-1]

    def getMin(self):
        """
        :rtype: int
        """
        return self.minVal


# Your MinStack object will be instantiated and called as such:
# obj = MinStack()
# obj.push(x)
# obj.pop()
# param_3 = obj.top()
# param_4 = obj.getMin()