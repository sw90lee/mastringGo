package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Tree노드 정의
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func traverse(t *Tree) {
	if t == nil {
		return
	}
	// 재귀호출로 방문
	traverse(t.Left)
	fmt.Print(t.Value, " ")
	traverse(t.Right)
}

// 이진트리를 정수 타입 난수로 채우는 함수
func create(n int) *Tree {
	var t *Tree
	rand.Seed(time.Now().Unix())
	for i := 0; i < 2*n; i++ {
		temp := rand.Intn(n * 2)
		t = insert(t, temp)
	}
	return t
}

func insert(t *Tree, v int) *Tree {
	//Tree가 비었는지 검사 -> 비어있다면,
	// 새로 생성한 노드가 트리루트가 되며 &Tree{nil,v,nil}로 생성
	if t == nil {
		return &Tree{nil, v, nil}
	}

	// Tree 추가하려는 값이 이미 트리에 들었는지 검사
	// 이미 존재한다면 아무 일도 하지 않고 t를 return
	if v == t.Value {
		return t
	}

	// 현재 노드의 상태에 따라 추가할 값을 노드의, 왼쪽에 넣을지 오른쪽에 넣을지 판단
	if v < t.Value {
		t.Left = insert(t.Left, v)
		return t
	}

	t.Right = insert(t.Right, v)
	return t
}

func main() {
	tree := create(10)
	fmt.Println("The value of the root of the tree is", tree.Value)
	traverse(tree)
	fmt.Println()
	tree = insert(tree, -10)
	tree = insert(tree, -2)
	traverse(tree)
	fmt.Println()
	fmt.Println("The value of the root of the tree is", tree.Value)
}
