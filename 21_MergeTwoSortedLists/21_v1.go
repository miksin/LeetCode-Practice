package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 0ms ??
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	var firstNode, l *ListNode
	if l1.Val < l2.Val {
		firstNode = l1
		l = l1
		l1 = l1.Next
	} else {
		firstNode = l2
		l = l2
		l2 = l2.Next
	}

	for l1 != nil || l2 != nil {
		if l1 == nil {
			l.Next = l2
			break
		} else if l2 == nil {
			l.Next = l1
			break
		} else if l1.Val < l2.Val {
			l.Next = l1
			l = l.Next
			l1 = l1.Next
		} else {
			l.Next = l2
			l = l.Next
			l2 = l2.Next
		}
	}

	return firstNode
}

func main() {
	a1, a2 := []int{1, 2, 4}, []int{1, 3, 4}
	l1, l2 := createListFromArray(a1), createListFromArray(a2)

	printList(l1)
	printList(l2)
	printList(mergeTwoLists(l1, l2))
}

func printList(start *ListNode) {
	for l := start; l != nil; l = l.Next {
		fmt.Printf("%d", l.Val)
		if l.Next != nil {
			fmt.Printf("->")
		}
	}
	fmt.Printf("\n")
}

func createListFromArray(arr []int) *ListNode {
	var firstNode *ListNode
	for i := len(arr) - 1; i >= 0; i-- {
		firstNode = &ListNode{Val: arr[i], Next: firstNode}
	}
	return firstNode
}
