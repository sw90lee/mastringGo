package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

// 원본행렬과 크기는 같지만 부호가 반대인 정수 원소로 리턴
func negativeMatrix(s [][]int) [][]int {
	for i, x := range s {
		for j, _ := range x {
			s[i][j] = -s[i][j]
		}
	}
	return s
}

// 두 행렬의 원소에 접근해 각각 더한 결과로 새로운 행렬을 만듦
func addMatrices(m1 [][]int, m2 [][]int) [][]int {
	result := make([][]int, len(m1))
	for i, x := range m1 {
		for j, _ := range x {
			result[i] = append(result[i], m1[i][j]+m2[i][j])
		}
	}
	return result
}

func main() {
	argument := os.Args
	if len(argument) != 3 {
		fmt.Println("Wrong Number of arguments!!")
		return
	}

	var row, col int
	//행
	row, err := strconv.Atoi(argument[1])
	if err != nil {
		fmt.Println("Need an interger: ", argument[1])
		return
	}
	//열
	col, err = strconv.Atoi(argument[2])
	if err != nil {
		fmt.Println("Need an interger: ", argument[2])
		return
	}
	fmt.Printf("Using %dx%d arrayes\n", row, col)

	// 양의 정수 검사
	if col <= 0 || row <= 0 {
		fmt.Println("Need positive matrix dimentsions!")
	}

	m1 := make([][]int, row)
	m2 := make([][]int, row)

	rand.Seed(time.Now().Unix())
	// m1과 m2를 난수로 초기화
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			m1[i] = append(m1[i], random(-1, i*j+rand.Intn(10)))
			m2[i] = append(m2[i], random(-1, i*j+rand.Intn(10)))
		}
	}

	fmt.Println("m1: ", m1)
	fmt.Println("m2: ", m2)

	// 덧셈
	r1 := addMatrices(m1, m2)

	// 뺄셈
	r2 := addMatrices(m1, negativeMatrix(m2))

	fmt.Println("r1: ", r1)
	fmt.Println("r2: ", r2)
}
