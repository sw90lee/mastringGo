package main

import (
	"fmt"
)

// HashTable의 Array사이즈
const ArraySize = 7

// Hastable sturct Array (사이즈가 7)
type HashTable struct {
	array [ArraySize]*bucket
}

// bucket struct (각 슬롯에 연결된 linked list)
type bucket struct {
	head *bucketNode
}

// bucket Node struct (각 bucket에 연결된 Node들)
type bucketNode struct {
	key  string
	next *bucketNode
}

// Insert ( 값을 넣고  hashTable에 배열 )
func (h *HashTable) Insert(key string) {
	// 넣고자하는 값의 ASCII으로 변환 화 array 사이즈로 나눈 값이 Index가 됨
	index := hash(key)
	h.array[index].bucketinsert(key)
}

// // // Search ( 값을 찾고 있다면 true 없다면 false )
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].bucketsearch(key)
}

// // // Delete ( key를 찾고 hashtable에서 삭제 )
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].bucketdelete(key)

}

// insert
func (b *bucket) bucketinsert(k string) {
	if !b.bucketsearch(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println(k, "이미 존재합니다.")
	}
}

// search
func (b *bucket) bucketsearch(k string) bool {
	currentNode := b.head
	// Node에 값이 비어있을때가지 찾음
	for currentNode != nil {
		// 현재의 노드 값이 k와 같다면 true
		if currentNode.key == k {
			return true
		}
		// 현재의 노드 값이 k와 아니면 if문없이 다음노드를 반환
		currentNode = currentNode.next
	}
	return false
}

// delete
func (b *bucket) bucketdelete(k string) {
	if b.head.key == k {
		b.head = b.head.next
		return
	}

	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == k {
			//delete
			previousNode.next = previousNode.next.next
		}
		previousNode = previousNode.next
	}
}

// hash
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

// init will create a bucket in each slot of the hash table
func Init() *HashTable {
	result := &HashTable{}
	// HashTable slot 즉 각 배열 내부에 bucket을 생성
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func main() {
	// HashTable 생성
	hashTable := Init()
	list := []string{
		"sungwoo",
		"misun",
		"simkong",
		"sungho",
		"soyun",
		"jaebok",
		"good",
	}

	for _, v := range list {
		hashTable.Insert(v)
	}

	hashTable.Delete("sungwoo")
}
