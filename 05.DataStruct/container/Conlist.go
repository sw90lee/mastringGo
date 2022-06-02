package main

import (
	"container/list"
	"fmt"
	"strconv"
)

func printList(l *list.List) {
	for t := l.Back(); t != nil; t = t.Prev() {
		fmt.Print(t.Value, " ")
	}
	fmt.Println()

	for t := l.Front(); t != nil; t = t.Next() {
		fmt.Print(t.Value, " ")
	}
	fmt.Println()
}

func main() {
	value := list.New()

	e1 := value.PushBack("One")
	e2 := value.PushBack("Two")
	value.PushFront("Three")
	value.InsertBefore("Four", e1)
	value.InsertAfter("Five", e2)
	value.Remove(e2)
	value.Remove(e2)
	value.InsertAfter("FiveFive", e2)
	value.PushBackList(value)

	printList(value)
	value.Init()

	fmt.Printf("After init(): %v\n", value)
	for i := 0; i < 20; i++ {
		value.PushFront(strconv.Itoa(i))
	}

	printList(value)
}
