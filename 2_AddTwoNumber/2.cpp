/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        int base = 10, carry = 0;
        ListNode *ans;
        ListNode *p1 = l1, *p2 = l2, **pa = &ans;
        
        while(p1 || p2){
            int d1 = p1? p1->val : 0;
            int d2 = p2? p2->val : 0;
            
            int digit = d1 + d2 + carry;
            carry = digit / base;
            digit %= base;
            
            *pa = new ListNode(digit);
            pa = &((*pa)->next);
            
            p1 = p1? p1->next : nullptr;
            p2 = p2? p2->next : nullptr;
        }
        
        if (carry != 0){
            *pa = new ListNode(carry);
        }
        
        return ans;
    }
};