package main

import (
	"fmt"
	"os"
	"reflect"
)

type t1 int
type t2 int

type a struct {
	X    int
	Y    float64
	Text string
}

func (a1 a) compareStruct(a2 a) bool {
	r1 := reflect.ValueOf(&a1).Elem()
	r2 := reflect.ValueOf(&a2).Elem()

	// numfield는 구조체안의 정의된 인스턴스를 알수있음
	for i := 0; i < r1.NumField(); i++ {
		// r1과 r2의 값을 비교후 틀리면 false
		if r1.Field(i).Interface() != r2.Field(i).Interface() {
			return false
		}
	}
	return true
}

// 메서드를 출력해주는 함수
func printMethods(i interface{}) {
	r := reflect.ValueOf(i)
	t := r.Type()
	fmt.Printf("type of examine: %s\n", t)

	// 메서드 안에 가진 갯수만큼 루프
	for j := 0; j < r.NumMethod(); j++ {
		m := r.Method(j).Type()
		fmt.Println(t.Method(j).Name, "--->", m)
	}
}

func main() {
	x1 := t1(100)
	x2 := t2(100)
	fmt.Printf("The type of x1 is %s\n", reflect.TypeOf(x1))
	fmt.Printf("The type of x2 is %s\n", reflect.TypeOf(x2))

	var p struct{}
	r := reflect.New(reflect.ValueOf(&p).Type()).Elem()
	fmt.Printf("The type of r is %s\n", reflect.TypeOf(r))

	a1 := a{1, 2.1, "A1"}
	a2 := a{1, -2, "A2"}

	if a1.compareStruct(a1) {
		fmt.Println("동일합니다.")
	}

	if !a1.compareStruct(a2) {
		fmt.Println("동일하지 않습니다.")
	}

	var f *os.File
	printMethods(f)
}
