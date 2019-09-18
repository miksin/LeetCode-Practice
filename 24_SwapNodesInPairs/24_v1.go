package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var p, pNext, pPrev *ListNode = head, head.Next, nil
	head = head.Next

	for {
		if pPrev != nil {
			pPrev.Next = pNext
		}

		p.Next = pNext.Next
		pNext.Next = p

		if p.Next == nil || p.Next.Next == nil {
			return head
		}

		pPrev = p
		p, pNext = p.Next, p.Next.Next
	}
}

func main() {
	input := createList([]int{1, 2, 3, 4})
	printNodes(input)
	printNodes(swapPairs(input))
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
