/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
public class Solution {
    public ListNode removeNthFromEnd(ListNode head, int n) {
        ArrayList<ListNode> l = new ArrayList<ListNode>();
        for (ListNode p = head; p != null; p = p.next)
            l.add(p);
            
        int idx = l.size() - n;
        ListNode removeNode = l.get(idx);
        
        if (idx == 0) {
            head = removeNode.next;
        } else {
            ListNode prevNode = l.get(idx - 1);
            prevNode.next = removeNode.next;
        }
        
        return head;
    }
}
