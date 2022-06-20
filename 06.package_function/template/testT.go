package main

import (
	"fmt"
	"os"
	"text/template"
)

type Entry struct {
	Number int
	Squere int
}

func main() {
	argument := os.Args
	if len(argument) != 2 {
		fmt.Println("Need The Template file!")
		return
	}

	tFile := argument[1]
	DATA := [][]int{{-1, 1}, {-2, 4}, {-3, 9}, {-4, 16}}
	var Entries []Entry
	for _, i := range DATA {
		if len(i) == 2 {
			temp := Entry{Number: i[0], Squere: i[1]}
			Entries = append(Entries, temp)
		}
	}

	// 초기화 작업 수행,  Template 값을 리턴
	// ParseGlob은 외부 템플릿 파일을 읽음
	t := template.Must(template.ParseGlob(tFile))
	t.Execute(os.Stdout, Entries)
}
