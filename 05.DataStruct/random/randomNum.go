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

func main() {
	// 출력 가능갯수 0~94개
	MIN := 0
	MAX := 94
	SEED := time.Now().Unix()
	// 출력시 문자 길이
	var LENGTH int64 = 8

	argument := os.Args

	switch len(argument) {
	case 2:
		LENGTH, _ = strconv.ParseInt(os.Args[1], 10, 64)
	default:
		fmt.Println("Using default value!!")
	}

	rand.Seed(SEED)
	// !은 10진수로 33임
	startChar := "!"
	var i int64 = 1
	for {
		myRand := random(MIN, MAX)
		newChar := string(startChar[0] + byte(myRand))
		fmt.Print(newChar)
		if i == LENGTH {
			break
		}
		i++
	}
	fmt.Println()
}
