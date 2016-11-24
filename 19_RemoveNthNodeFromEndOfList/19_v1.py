# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def removeNthFromEnd(self, head, n):
        """
        :type head: ListNode
        :type n: int
        :rtype: ListNode
        """
        l = []
        p = head
        while p is not None:
            l.append(p)
            p = p.next
        
        idx = len(l) - n
        
        if idx == 0:
            head = l[idx].next
        else:
            l[idx-1].next = l[idx].next
            
        return head
        
