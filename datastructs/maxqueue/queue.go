package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	items []string
	head  int
	tail  int
	size  int
	limit int
}

func NewQueue(length int) *Queue {
	return &Queue{
		items: make([]string, length),
		head:  0,
		tail:  0,
		size:  0,
		limit: length,
	}
}

func (q *Queue) Push(item string) {
	if q.size < q.limit {
		q.items[q.tail] = item
		q.tail = (q.tail + 1) % q.limit
		q.size = q.size + 1
	}
}

func (q *Queue) Pop() (string, error) {
	if q.IsEmpty() {
		return "", errors.New("cannot pop element from empty queue")
	}

	item := q.items[q.head]
	q.items[q.head] = ""
	q.head = (q.head + 1) % q.limit
	q.size = q.size - 1

	return item, nil
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func main() {
	q := NewQueue(8)
	q.Push("1")
	fmt.Println(q.items, q.size)
	q.Push("-1")
	q.Push("0")
	q.Push("10")
	q.Pop()
	fmt.Println(q.items, q.size)
}
