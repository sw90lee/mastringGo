package main

import "fmt"

func main() {
	iMap := map[string]int{
		"k1":12,
		"k2":13,
	}
	fmt.Println("imap:",iMap)

	_,ok := iMap["doesItExist"]
}
