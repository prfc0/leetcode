package main

/*
https://leetcode.com/problems/merge-two-sorted-lists/
You are given the heads of two sorted linked lists list1 and list2.
Merge the two lists in a one sorted list. The list should be made by splicing together the nodes of the first two lists.
Return the head of the merged linked list.
*/

import "fmt"

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	dummy := &ListNode{Val: 0}

	if list1 == nil && list2 == nil {
		return dummy.Next
	}

	curr := dummy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			curr.Next = list1
			curr = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			curr = list2
			list2 = list2.Next
		}
	}
	for list1 != nil {
		curr.Next = list1
		curr = list1
		list1 = list1.Next
	}
	for list2 != nil {
		curr.Next = list2
		curr = list2
		list2 = list2.Next
	}
	return dummy.Next
}

func main() {
	fmt.Println("vim-go")
}
