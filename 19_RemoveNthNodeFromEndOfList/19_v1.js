/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} head
 * @param {number} n
 * @return {ListNode}
 */
var removeNthFromEnd = function(head, n) {
    var l = []
    for (let p = head; p !== null; p = p.next){
        l.push(p)
    }
    
    var idx = l.length - n
    if (idx === 0){
        head = l[idx].next
    } else {
        l[idx - 1].next = l[idx].next
    }
    
    return head
};
