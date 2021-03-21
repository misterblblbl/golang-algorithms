package datastructs

import (
	"errors"
	"fmt"
)

type ListNode struct {
	next  *ListNode
	prev  *ListNode
	value int
}

func (s *Stack) PrintStack() {
	node := s.head
	for node != nil {
		fmt.Print("*", node.value, " ")
		node = node.prev
	}
}

type Stack struct {
	head *ListNode
	max  int
}

func (s *Stack) Push(item int) {
	newNode := &ListNode{value: item, prev: nil}

	if s.head == nil {
		s.head = newNode
	} else {
		newNode.prev = s.head
		s.head = newNode
	}
}

func (s *Stack) Pop() (int, error) {
	if s.head == nil {
		return 0, errors.New("stack is empty")
	}

	value := s.head.value
	s.head = s.head.prev

	return value, nil
}

func (s *Stack) GetMax() int {
	if s.head == nil {
		fmt.Println("None")
		return 0
	}

	node := s.head
	max := s.head.value

	for node != nil {
		if node.value > max {
			max = node.value
		}

		node = node.prev
	}
	fmt.Println(max)
	return max
}

// func (s *Stack) GetMax() int {
// 	// should show the max item
// }

func RunExampleStack() {
	stack := &Stack{}
	stack.GetMax()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(-3)
	stack.GetMax()

	stack.PrintStack()

	item, _ := stack.Pop()
	fmt.Println("IR: ", item)
	stack.Push(-3)
	stack.PrintStack()
}
