package main

import "fmt"

type Queue struct {
	item []int
}

func (q *Queue) Enqueue(i int) {
	q.item = append(q.item, i)
}

func (q *Queue) DeQueue() int {
	toRemove := q.item[0]
	q.item = q.item[1:]
	return toRemove
}

func main() {
	myQueue := Queue{}
	fmt.Println(myQueue)
	myQueue.Enqueue(100)
	myQueue.Enqueue(200)
	myQueue.Enqueue(300)
	fmt.Println(myQueue)
	myQueue.DeQueue()
	fmt.Println(myQueue)
}
