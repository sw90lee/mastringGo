package main

import (
	"fmt"
)

var count int

type Node struct {
	key   int
	left  *Node
	Right *Node
}

//Insert
func (n *Node) Insert(k int) {
	if n.key < k {
		//move right
		if n.Right == nil {
			n.Right = &Node{key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.key > k {
		//move left
		if n.left == nil {
			n.left = &Node{key: k}
		} else {
			n.left.Insert(k)
		}
	}

}

func (n *Node) Search(k int) bool {

	count++

	if n == nil {
		return false
	}

	if n.key < k {
		//move right
		return n.Right.Search(k)
	} else if n.key > k {
		//move left
		return n.left.Search(k)
	}

	// k의 값이 만약에 찾는 위치에 있다면 true 반환
	return true

}

func main() {
	tree := &Node{key: 100}
	fmt.Println(tree)
	tree.Insert(52)
	tree.Insert(32)
	tree.Insert(252)
	tree.Insert(552)
	tree.Insert(512)
	tree.Insert(92)

	fmt.Println(tree.Search(100))
	fmt.Println(count)
}
