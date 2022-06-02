package main

import (
	"fmt"
)

type Stack struct {
	item []int
}

// Push
func (s *Stack) Push(i int) {
	s.item = append(s.item, i)
}

// Pop

func (s *Stack) Pop() {
	l := len(s.item) - 1
	s.item = s.item[:l]
	return
}

func main() {
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)
	myStack.Push(400)
	fmt.Println(myStack)
	myStack.Pop()
	myStack.Pop()
	fmt.Println(myStack)
	myStack.Push(500)
	fmt.Println(myStack)
}
