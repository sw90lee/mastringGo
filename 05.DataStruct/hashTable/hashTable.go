package main

import "fmt"

const SIZE = 15

type Node struct {
	Value int
	Next  *Node
}

// Table 연결리스트에 대한 맵
// Table 크기
type HashTable struct {
	Table map[int]*Node
	Size  int
}

//
func hashFucntion(i, size int) int {
	return (i % size)
}

// 해쉬 테이블 원소 추가할때 호출
func insert(hash *HashTable, value int) int {
	index := hashFucntion(value, hash.Size)
	element := Node{Value: value, Next: hash.Table[index]}
	hash.Table[index] = &element
	return index
}

// 해시 테이블에 담긴 모든 값을 출력
// 해시 테이블에 있는 연결 리스트에 담긴 모든 원소를 방문해 그값을 화면에 출력함
// 연결리스트 단위로 처리
func traverse(hash *HashTable) {
	for k := range hash.Table {
		if hash.Table[k] != nil {
			t := hash.Table[k]
			for t != nil {
				fmt.Printf("%d -> ", t.Value)
				t = t.Next
			}
			fmt.Println()
		}
	}
}

func lookup(hash *HashTable, value int) bool {
	index := hashFucntion(value, hash.Size)
	if hash.Table[index] != nil {
		t := hash.Table[index]
		for t != nil {
			if t.Value == value {
				fmt.Printf("%d는 있습니다.\n", t.Value)
				return true
			}
			t = t.Next
		}
	}
	return false
}

func main() {
	// 해쉬 테이블 생성 (SIZE 15)
	table := make(map[int]*Node, SIZE)

	hash := &HashTable{Table: table, Size: SIZE}
	fmt.Println("Number of space:", hash.Size)
	for i := 0; i < 120; i++ {
		insert(hash, i)
	}
	traverse(hash)

	fmt.Println(lookup(hash, 1))
}
