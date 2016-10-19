/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
public class Solution {
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        int base = 10;
        int carry = 0;
        ListNode ans = new ListNode(0), pa = ans;

        if (l1 != null || l2 != null){
            int d1 = (l1 != null)? l1.val: 0;
            int d2 = (l2 != null)? l2.val: 0;

            int digit = d1 + d2 + carry;
            carry = digit / base;
            digit %= base;

            ans = new ListNode(digit);
            pa = ans;
            
            l1 = (l1 != null)? l1.next: null;
            l2 = (l2 != null)? l2.next: null;
        }
        
        while (l1 != null || l2 != null){
            int d1 = (l1 != null)? l1.val: 0;
            int d2 = (l2 != null)? l2.val: 0;

            int digit = d1 + d2 + carry;
            carry = digit / base;
            digit %= base;

            pa.next = new ListNode(digit);
            pa = pa.next;
            
            l1 = (l1 != null)? l1.next: null;
            l2 = (l2 != null)? l2.next: null;
        }
        
        if (carry != 0){
            pa.next = new ListNode(carry);
        }

        return ans;
    }
}