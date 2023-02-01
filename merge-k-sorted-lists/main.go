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

func createLinkedList(nums []int) *ListNode {
	var head *ListNode = nil
	n := len(nums)
	for i := n - 1; i >= 0; i-- {
		nn := &ListNode{Val: nums[i]}
		nn.Next = head
		head = nn
	}
	return head
}

func printLinkedList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " -> ")
		head = head.Next
	}
	fmt.Println()
}

type LLHeap []*ListNode

func (h LLHeap) Len() int           { return len(h) }
func (h LLHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h LLHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *LLHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *LLHeap) Pop() interface{} {
	old := *h
	N := len(old)
	x := old[N-1]
	*h = old[:N-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := &LLHeap{}
	heap.Init(h)
	for _, head := range lists {
		if head != nil {
			heap.Push(h, head)
		}
	}

	dummy := &ListNode{}
	curr := dummy
	for h.Len() > 0 {
		node := heap.Pop(h).(*ListNode)
		curr.Next = node
		if node.Next != nil {
			heap.Push(h, node.Next)
		}
		curr = node
	}
	return dummy.Next
}

type Input struct {
	lists [][]int
}

func main() {
	inputs := []Input{
		{
			lists: [][]int{
				{1, 4, 5},
				{1, 3, 4},
				{2, 6},
			},
		},
		{
			lists: [][]int{},
		},
		{
			lists: [][]int{
				{},
			},
		},
	}
	for _, input := range inputs {
		heads := make([]*ListNode, len(input.lists))
		for i, list := range input.lists {
			head := createLinkedList(list)
			printLinkedList(head)
			heads[i] = head
		}
		head := mergeKLists(heads)
		fmt.Print("Merged list: ")
		printLinkedList(head)
		fmt.Println()
	}
}
