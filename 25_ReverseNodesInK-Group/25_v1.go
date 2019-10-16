package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k < 2 || head == nil {
		return head
	}

	conn := &head
	start, end := head, head
	curr := head
	var prev *ListNode
	i := 0

	for curr != nil {
		i++

		next := curr.Next
		curr.Next = prev
		prev = curr
		end = curr
		curr = next

		if i == k {
			*conn = end
			conn = &start.Next
			start.Next = next
			start = next
			i = 0
		}
	}

	if i < k {
		prev, curr = curr, prev
		for i > 0 {
			next := curr.Next
			curr.Next = prev
			prev = curr
			end = curr
			curr = next
			i--
		}
	}

	return head
}

func main() {
	printNodes(createList([]int{1, 2, 3, 4, 5}))
	printNodes(reverseKGroup(createList([]int{1, 2, 3, 4, 5}), 2))
	printNodes(reverseKGroup(createList([]int{1, 2, 3, 4, 5}), 3))
	printNodes(reverseKGroup(createList([]int{1, 2, 3, 4, 5}), 4))
	printNodes(reverseKGroup(createList([]int{1, 2, 3, 4, 5}), 5))
}

func printNodes(p *ListNode) {
	for p != nil {
		fmt.Printf("%d", p.Val)
		if p.Next != nil {
			fmt.Printf("->")
		}
		p = p.Next
	}
	fmt.Println("")
}

func createList(arr []int) *ListNode {
	var p *ListNode
	for i := len(arr) - 1; i >= 0; i-- {
		node := ListNode{
			Val:  arr[i],
			Next: p,
		}
		p = &node
	}
	return p
}
