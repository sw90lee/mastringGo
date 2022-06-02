package main

import (
	"container/ring"
	"fmt"
)

var size int = 10

func main() {
	// 새로운 링 생성
	myRing := ring.New(size + 1)
	// Empty ring: {0xc0000a43c0 0xc0000a44e0 <nil>}
	fmt.Println("Empty ring:", *myRing)
	for i := 0; i < myRing.Len()-1; i++ {
		myRing.Value = i
		myRing = myRing.Next()
	}
	// 마지막 Ring의 값 지정을 안하면 nil의 값이지만 넣어주면
	// Ring의 마지막 즉 myRing.Value가 나옴
	// Empty ring: {0xc0000a43c0 0xc0000a44e0 3}
	myRing.Value = 3
	sum := 0
	// 각 함수 호출시 사용
	// 모든 원소 순환에 사용에는 ring.Do()함수로 하는것이 선호
	myRing.Do(func(x interface{}) {
		// type assertion
		t := x.(int)
		sum = sum + t
	})
	fmt.Println("Sum:", sum)

	for i := 0; i < myRing.Len()+3; i++ {
		myRing = myRing.Next()
		fmt.Print(myRing.Value, " ")
	}
	fmt.Println()
	fmt.Println("Empty ring:", *myRing)
}
