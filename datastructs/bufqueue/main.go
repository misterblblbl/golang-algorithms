package main

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
)

type BufferedQueue struct {
	buffer []*int
	max    int
	size   int
	head   int
	tail   int
}

func NewBufferedQueue(bufSize int) *BufferedQueue {
	buf := make([]*int, bufSize)

	return &BufferedQueue{
		buffer: buf,
		size:   0,
		head:   0,
		tail:   0,
		max:    bufSize,
	}
}

func (q *BufferedQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *BufferedQueue) Enqueue(item int) {
	if q.size < q.max {
		q.buffer[q.tail] = &item
		q.tail = (q.tail + 1) % q.max
		q.size = q.size + 1
	}
}

func (q *BufferedQueue) Dequeue() (*int, error) {
	if q.IsEmpty() {
		return nil, errors.New("the queue is empty")
	}

	item := q.buffer[q.head]
	q.buffer[q.head] = nil
	q.head = (q.head + 1) % q.max
	q.size = q.size - 1

	return item, nil
}

func (q *BufferedQueue) String() string {
	var buffer bytes.Buffer

	for _, item := range q.buffer {
		if item != nil {
			buffer.WriteString(strconv.Itoa(*item))
			buffer.WriteString(" ")
		}
	}

	return "[" + buffer.String() + "]"
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func getSlidingMinimum(data []int, window int) []int {
	mins := make([]int, len(data))
	q := NewBufferedQueue(window)

	lastMinimum, minIndex := 0, 0

	for i, item := range data[:window] {
		if i == 0 {
			mins[0] = item
		} else {
			if mins[0] > item {
				lastMinimum = i
				mins[0] = item
			}
		}

		q.Enqueue(item)
	}

	fmt.Println("lastMinimum", lastMinimum)

	for i := 1; i < len(data); i++ {
		item := data[i]
		q.Enqueue(item)

		withinRange := lastMinimum >= i && lastMinimum <= i+(window-1)
		fmt.Println("minIndex===", i, lastMinimum, withinRange, i, i+(window-1))
		if !withinRange {
			lastMinimum = i
		}

		if data[lastMinimum] > item {
			lastMinimum = i
			mins[minIndex+1] = item
		} else {
			mins[minIndex+1] = data[lastMinimum]
		}

		minIndex++
	}

	return mins
}

func main() {
	arr := []int{5, 1, 3, 4, 6, 1, 7, 3, 2, 8}
	mins := getSlidingMinimum(arr, 3)
	// q := NewBufferedQueue(3)
	// q.Enqueue(1)
	// q.Enqueue(2)
	// q.Enqueue(3)

	// fmt.Println(q.buffer)

	// x, _ := q.Dequeue()

	fmt.Println(mins)
}
