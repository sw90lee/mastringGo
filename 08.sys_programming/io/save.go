package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	s := []byte("Data to write\n")

	// f1 생성
	f1, err := os.Create("f1.txt")
	if err != nil {
		fmt.Println("Cannot create file", err)
		return
	}
	defer f1.Close()
	fmt.Fprintf(f1, string(s))

	// f2 생성
	f2, err := os.Create("f2.txt")
	if err != nil {
		fmt.Println("Cannot create file", err)
		return
	}
	defer f2.Close()
	n, err := f2.WriteString(string(s))
	fmt.Printf("wrote %d bytes\n", n)

	f3, err := os.Create("f3.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// f3 생성
	w := bufio.NewWriter(f3)
	n, err = w.WriteString(string(s))
	fmt.Printf("wrote %d byte\n", n)
	w.Flush()

	// f4 생성
	f4 := "f4.txt"
	err = ioutil.WriteFile(f4, s, 0644)
	if err != nil {
		fmt.Println("Cannot create file", err)
		return
	}

	// f5 생성
	f5, err := os.Create("f5.txt")
	if err != nil {
		fmt.Println("Cannot create file", err)
		return
	}
	n, err = io.WriteString(f5, string(s))
	if err != nil {
		fmt.Println("Cannot create file", err)
		return
	}
	fmt.Printf("wrote %d byte\n", n)

}
