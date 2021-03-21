package datastructs

import "fmt"

type ListNode struct {
	next  *ListNode
	prev  *ListNode
	value string
}

func (n *ListNode) String() string {
	if n != nil {
		return n.value
	}

	return "-"
}

func WalkList(head *ListNode) {
	for head != nil {
		fmt.Println(head.value)
		head = head.next
	}
}

func WalkListReversed(head *ListNode) {
	for head != nil {
		fmt.Println(head.value)
		head = head.prev
	}
}

func ReverseList(node *ListNode) *ListNode {
	for node.next != nil {
		fmt.Println("VALUE:", node, "  NEXT:", node.next, "  PREV:", node.prev)
		node.next, node.prev = node.prev, node.next
		node = node.prev
	}

	node.next, node.prev = node.prev, node.next

	fmt.Println("FIN:", node)

	return node
}

func findByIndex(head *ListNode, index int) *ListNode {
	for index > 0 && head.next != nil {
		head = head.next
		index = index - 1
	}

	return head
}

func InsertIntoList(head *ListNode, index int, value string) *ListNode {
	newNode := &ListNode{value: value}

	if index == 0 {
		newNode.next = head
		return newNode
	}

	previuosNode := findByIndex(head, index-1)
	newNode.next = previuosNode.next
	previuosNode.next = newNode

	return head
}

func DeleteFromList(head *ListNode, index int) *ListNode {
	if index == 0 {
		return head.next
	}

	previuosNode := findByIndex(head, index-1)
	currNode := previuosNode.next
	previuosNode.next = currNode.next

	return head
}

func FindInList(node *ListNode, value string) int {
	var index int

	for node != nil {
		if node.value == value {
			return index
		}

		node = node.next
		index++
	}

	return -1
}

func RunExample() {
	// n3 := &ListNode{value: "three", next: nil, prev: nil}
	// n2 := &ListNode{value: "two", next: n3, prev: nil}
	// n1 := &ListNode{value: "one", next: n2, prev: nil}
	// WalkList(n1)

	d1 := &ListNode{value: "one", next: nil, prev: nil}
	d2 := &ListNode{value: "two", next: nil, prev: d1}
	// d3 := &ListNode{value: "three", next: nil, prev: d2}
	d1.next = d2
	// d2.next = d3

	WalkList(d1)

	fmt.Println("==========REVERSED========")
	rev := ReverseList(d1)
	WalkList(rev)
}
