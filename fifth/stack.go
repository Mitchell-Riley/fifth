package fifth

import (
	"fmt"
	"log"
)

type stack struct {
	top *node
}

type node struct {
	data int
	next *node
}

// print stack so it can be read left to right with the bottom of the
// stack getting printed first
func (s stack) String() string {
	reversedStack := new(stack)
	for current := s.top; current != nil; current = current.next {
		reversedStack.Push(current)
	}

	var str string
	for current := reversedStack.top; current != nil; current = current.next {
		str += fmt.Sprint(current) + " "
	}

	return str
}

func (n node) String() string {
	return fmt.Sprint(n.data)
}

func (s *stack) isEmpty() bool {
	return s.top == nil
}

func (s *stack) size() int {
	size := 0
	for current := s.top; current != nil; current = current.next {
		size++
	}
	return size
}

func (s *stack) search(idx int) *node {
	loopCount := 0
	current := s.top
	for current != nil {
		if loopCount == idx {
			return current
		}
		loopCount++
		current = current.next
	}
	return nil
}

func (s *stack) Push(data interface{}) {
	n := s.top
	switch data.(type) {
	case int:
		s.top = &node{data: data.(int)}
	case int32:
		s.top = &node{data: int(data.(int32))}
	case *node:
		s.top = &node{data: data.(*node).data}
	default:
		log.Fatalf("Unknown type %T", data)
	}
	s.top.next = n
}

func (s *stack) Pop() *node {
	if s.isEmpty() {
		return nil
	}

	n := s.top
	s.top = s.top.next
	return n
}
