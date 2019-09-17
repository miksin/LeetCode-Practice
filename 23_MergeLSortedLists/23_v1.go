package main

import (
	"container/heap"
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

type PriorityQueue []*ListNode

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Val < pq[j].Val
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	node := x.(*ListNode)
	*pq = append(*pq, node)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return node
}

func mergeKLists(lists []*ListNode) *ListNode {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	for _, node := range lists {
		if node != nil {
			heap.Push(&pq, node)
		}
	}

	var root, current *ListNode

	if pq.Len() < 1 {
		return nil
	} else {
		node := heap.Pop(&pq).(*ListNode)
		next := node.Next
		node.Next = nil

		current = node
		root = node

		if next != nil {
			heap.Push(&pq, next)
		}
	}

	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*ListNode)
		next := node.Next
		node.Next = nil

		current.Next = node
		current = current.Next

		if next != nil {
			heap.Push(&pq, next)
		}
	}

	return root
}

func main() {
	input := []*ListNode{
		createList([]int{1, 4, 5}),
		createList([]int{1, 3, 4}),
		createList([]int{2, 6}),
	}

	printNodes(input[0])
	printNodes(input[1])
	printNodes(input[2])
	printNodes(mergeKLists(input))
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

/*
Input:
[
  1->4->5,
  1->3->4,
  2->6
]
Output: 1->1->2->3->4->4->5->6
*/
