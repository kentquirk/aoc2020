package main

import (
	"fmt"
	"strings"
)

// Element is a single item in our doubly-linked list
type Element struct {
	next  *Element
	value int
}

// Next returns the nth next element
// Next(0) is the element itself
func (e *Element) Next(n int) *Element {
	for n > 0 {
		e = e.next
		n--
	}
	return e
}

// List is a doubly-linked list; head points to the "current"
// element, and side is used to keep track of a side list
type List struct {
	head     *Element
	maxlabel int
	lookup   map[int]*Element
}

func newList(data []int, maxlabel int) *List {
	lookup := make(map[int]*Element)
	first := &Element{value: data[0]}
	lookup[data[0]] = first
	prev := first
	for _, v := range data[1:] {
		elt := &Element{next: nil, value: v}
		prev.next = elt
		prev = elt
		lookup[v] = elt
	}

	for v := 10; v <= maxlabel; v++ {
		elt := &Element{next: nil, value: v}
		prev.next = elt
		prev = elt
		lookup[v] = elt
	}

	prev.next = first

	return &List{
		head:     first,
		maxlabel: maxlabel,
		lookup:   lookup,
	}
}

// Find locates an element with the given value, or
// returns nil if not found
func (l *List) Find(v int) *Element {
	if l.lookup != nil {
		return l.lookup[v]
	}
	if l.head.value == v {
		return l.head
	}
	for p := l.head.next; p != l.head; p = p.next {
		if p.value == v {
			return p
		}
	}
	return nil
}

// FindDestination locates the destination cup and returns a pointer to it.
func (l *List) FindDestination(excludes *List) *Element {
	label := l.head.value - 1
	for {
		if label < 1 {
			label = l.maxlabel
		}
		if excludes.Find(label) == nil {
			break
		}
		label--
	}

	return l.Find(label)
}

// Move relocates a set of cups according to the rules
func (l *List) Move() {
	// move 3 cups after the head to the side
	side := &List{head: l.head.next}
	l.head.next = l.head.Next(4)
	side.head.Next(2).next = side.head
	// find the place to put them
	dest := l.FindDestination(side)
	// put the side cups back
	side.head.Next(2).next = dest.next
	dest.next = side.head
	// and re-aim the head
	l.head = l.head.next
}

func (l *List) String() string {
	sb := &strings.Builder{}
	fmt.Fprintf(sb, "(%d) ", l.head.value)
	for p := l.head.next; p != l.head; p = p.next {
		fmt.Fprintf(sb, "%d ", p.value)
	}
	return sb.String()
}

// ResultA prints the result in the right format
func (l *List) ResultA() string {
	sb := &strings.Builder{}
	one := l.Find(1)
	for p := one.next; p != one; p = p.next {
		fmt.Fprintf(sb, "%d", p.value)
	}
	return sb.String()
}

// ResultB prints the result for part b in the right format.
func (l *List) ResultB() string {
	sb := &strings.Builder{}
	one := l.Find(1)
	v1 := one.Next(1).value
	v2 := one.Next(2).value
	prod := v1 * v2
	fmt.Fprintf(sb, "%d*%d == %d", v1, v2, prod)
	return sb.String()
}

func day23a(data []int) string {
	list := newList(data, 9)
	fmt.Println(list.String())
	for i := 0; i < 100; i++ {
		list.Move()
		fmt.Println(list.String())
	}
	return list.ResultA()
}

func day23b(data []int) string {
	list := newList(data, 1_000_000)
	// fmt.Println(list.String())
	for i := 0; i < 10_000_000; i++ {
		list.Move()
		// fmt.Println(list.String())
	}
	return list.ResultB()
}

func main() {
	// data := []int{3, 8, 9, 1, 2, 5, 4, 6, 7} // sample
	data := []int{3, 8, 9, 5, 4, 7, 6, 1, 2} // real
	// fmt.Println(day23a(data))
	fmt.Println(day23b(data))
}
