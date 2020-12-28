package main

import (
	"log"
	"strconv"
	"strings"
)

// IntQueue manages a queue of integers that is capped at a
// maximum size. It doesn't deal well with errors.
// Tail points past the last item in the queue
// Head points to the next item in the queue
// If head and tail are equal, the queue is empty
type IntQueue struct {
	store    []int
	capacity int
	head     int
	tail     int
}

func newIntQueue(capacity int) *IntQueue {
	return &IntQueue{
		store:    make([]int, capacity),
		capacity: capacity,
		tail:     0,
		head:     0,
	}
}

func (q *IntQueue) increment(n int) int {
	n++
	if n >= q.capacity {
		n = 0
	}
	return n
}

func (q *IntQueue) decrement(n int) int {
	n--
	if n < 0 {
		n = q.capacity - 1
	}
	return n
}

// Enqueue adds an integer to the queue
func (q *IntQueue) Enqueue(v int) {
	newtail := q.increment(q.tail)
	if newtail == q.head {
		log.Fatal("queue is full!")
	}
	q.store[q.tail] = v
	q.tail = newtail
}

// Dequeue removes an item from the queue and returns it
func (q *IntQueue) Dequeue() int {
	if q.tail == q.head {
		log.Fatal("queue is empty!")
	}
	v := q.store[q.head]
	q.head = q.increment(q.head)
	return v
}

// Size returns the number of items in the queue
func (q *IntQueue) Size() int {
	if q.tail >= q.head {
		return q.tail - q.head
	}
	// queue has wrapped around
	return q.capacity - (q.head - q.tail)
}

// Clone duplicates a queue exactly
func (q *IntQueue) Clone() *IntQueue {
	r := &IntQueue{
		store:    make([]int, q.capacity),
		capacity: q.capacity,
		tail:     q.tail,
		head:     q.head,
	}
	copy(r.store, q.store)
	return r
}

// Subset duplicates a queue with only a subset of the contents
// It throws away items from the tail until the queue is the right size.
func (q *IntQueue) Subset(qty int) *IntQueue {
	r := q.Clone()
	for ; r.Size() > qty; r.tail = r.decrement(r.tail) {
	}
	return r
}

// Hash generates a hash from the contents of the queue
// I'm doing a lame integer hash routine
func (q *IntQueue) Hash() int {
	x := 0x45d9f3b
	for p := q.head; p != q.tail; p = q.increment(p) {
		x ^= ((q.store[p] >> 16) ^ q.store[p]) * 0x45d9f3b
		x = ((x >> 16) ^ x) * 0x45d9f3b
		x = (x >> 16) ^ x
	}
	return x
}

func (q *IntQueue) String() string {
	s := []string{}
	for p := q.head; p != q.tail; p = q.increment(p) {
		s = append(s, strconv.Itoa(q.store[p]))
	}
	return strings.Join(s, ", ")
}
