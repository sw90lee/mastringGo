package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

var size = 0
var queue = new(Node)

func Push(t *Node, value int) bool {
	if queue == nil {
		queue = &Node{Value: value}
		size++
		return true
	}

	t = &Node{value, nil}
	t.Next = queue
	queue = t
	size++
	return true
}

func Pop(t *Node) (int, bool) {
	if size == 0 {
		return 0, false
	}

	if size == 1 {
		queue = nil
		size--
		return t.Value, true
	}

	temp := t
	for t.Next != nil {
		temp = t
		t = t.Next
	}

	v := temp.Next.Value
	temp.Next = nil

	size--
	return v, true
}

func traverse(t *Node) {
	if size == 0 {
		fmt.Println("Empty Queue!!")
		return
	}

	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
}

func main() {
	queue = nil
	Push(queue, 10)
	fmt.Println("Size: ", size)
	Push(queue, 100)
	fmt.Println()
	Push(queue, 5)
	traverse(queue)
	fmt.Println(Pop(queue))
	traverse(queue)

}
