package main

/*
https://leetcode.com/problems/valid-parentheses/
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
An input string is valid if:
Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
*/

import "fmt"

type node struct {
	value byte
	prev  *node
}

type Stack struct {
	top    *node
	length int
}

func NewStack() *Stack {
	return &Stack{nil, 0}
}

func (s *Stack) Len() int {
	return s.length
}

func (s *Stack) Peek() byte {
	if s.length == 0 {
		panic("Cannot peek an empty stack.")
	}
	return s.top.value
}

func (s *Stack) Pop() byte {
	if s.length == 0 {
		panic("Cannot pop from empty stack.")
	}
	n := s.top
	s.top = n.prev
	s.length--
	return n.value
}

func (s *Stack) Push(value byte) {
	n := &node{value, s.top}
	s.top = n
	s.length++
}

func isValid(s string) bool {
	st := NewStack()
	for i := 0; i < len(s); i++ {
		p := s[i]
		if p == '(' || p == '[' || p == '{' {
			st.Push(s[i])
		} else if p == ')' || p == ']' || p == '}' {
			if st.Len() == 0 {
				return false
			}
			popped := st.Pop()
			if p == ')' && popped != '(' {
				return false
			}
			if p == ']' && popped != '[' {
				return false
			}
			if p == '}' && popped != '{' {
				return false
			}
		}
	}
	if st.Len() != 0 {
		return false
	}
	return true
}

func main() {
	// P := "()"
	P := "([])[((()))]{}"
	// P := "(]"
	fmt.Println(isValid(P))
}
