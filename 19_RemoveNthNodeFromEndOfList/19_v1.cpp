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
    ListNode* removeNthFromEnd(ListNode* head, int n) {
        vector<ListNode*> v;
        for (ListNode* p = head; p != nullptr; p = p->next)
            v.push_back(p);
        
        int idx = v.size() - n;
        ListNode *removeNode = v[idx];
        
        if (idx == 0){
            head = removeNode->next;
        } else {
            ListNode *prevNode = v[idx - 1];
            prevNode->next = removeNode->next;
        }        

        return head;
    }
};
