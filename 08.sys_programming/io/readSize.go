package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func readSize(f *os.File, size int) []byte {
	buffer := make([]byte, size)

	n, err := f.Read(buffer)
	if err == io.EOF {
		return nil
	}

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return buffer[0:n]
}

func main() {
	argument := os.Args
	if len(argument) != 3 {
		fmt.Println("<buffer size> <filename>")
		return
	}

	// Int 반환
	bufferSize, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	file := os.Args[2]
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	for {
		readData := readSize(f, bufferSize)
		if readData != nil {
			fmt.Print(string(readData))
		} else {
			break
		}
	}
}
