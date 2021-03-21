package datastructs

import (
	"fmt"
	"strings"
)

// Stack is stack based on an array

type Stack struct {
	items []string
}

func NewStack() *Stack {
	return &Stack{items: make([]string, 0, 100)}
}

func (s *Stack) String() string {
	return "[" + strings.Join(s.items, ", ") + "]"
}

func (s *Stack) Push(item string) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() string {
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return item
}

func (s *Stack) Peek() string {
	return s.items[len(s.items)-1]
}

func (s *Stack) Size() int {
	return len(s.items)
}

//
// StackEffective is stack based on a linked list

func RunExampleStack() {
	stack := NewStack()
	stack.Push("apple")
	stack.Push("banana")
	stack.Push("orange")

	fmt.Println(stack)

	stack.Pop()
	fmt.Println(stack)
}
