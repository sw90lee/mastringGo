package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

// 연결리스트의 첫번째 원소 즉 root
var root = new(Node)

func addNode(t *Node, v int) int {
	// 연결리스트 비어있는지 확인
	if root == nil {
		t = &Node{v, nil}
		root = t
		return 0
	}

	// 추가하려는 값 중복 확인
	if v == t.Value {
		fmt.Println("Node already exist : ", v)
		return -1
	}

	// 노드의 끝 확인
	if t.Next == nil {
		t.Next = &Node{v, nil}
		return -2
	}

	// 3가지에 걸리지 않으면 다음 노드에 대해 동작 반복
	return addNode(t.Next, v)
}

func traverse(t *Node) {
	if t == nil {
		fmt.Println("-> Empty List!!!")
		return
	}

	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}

	fmt.Println()
}

// 주어진 원소가 연결리스트에 존재하는지 확인
func lookupNode(t *Node, v int) bool {
	if root == nil {
		t = &Node{v, nil}
		root = t
		return false
	}

	if v == t.Value {
		return true
	}

	if t.Next == nil {
		return false
	}

	return lookupNode(t.Next, v)
}

// 연결리스트 크기
func size(t *Node) int {
	if t == nil {
		fmt.Println("-> Empty list!!!")
		return 0
	}

	i := 0
	for t != nil {
		i++
		t = t.Next
	}
	return i
}

func main() {
	fmt.Println(root)
	root = nil
	traverse(root)
	addNode(root, 1)
	addNode(root, -1)
	traverse(root)
	addNode(root, 10)
	addNode(root, 5)
	addNode(root, 45)
	traverse(root)

	addNode(root, 100)
	traverse(root)

	if lookupNode(root, 100) {
		fmt.Println("Node exist!!")
	} else {
		fmt.Println("Node does not exist!!")
	}

	if lookupNode(root, -100) {
		fmt.Println("Node exist!!")
	} else {
		fmt.Println("Node does not exist!!")
	}

}
