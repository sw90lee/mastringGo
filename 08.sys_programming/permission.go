package main

import (
	"fmt"
	"os"
)

func main() {
	argument := os.Args
	if len(argument) == 1 {
		fmt.Println("usage: permission filename\n")
		return
	}

	filename := argument[1]
	info, _ := os.Stat(filename)
	mode := info.Mode()
	fmt.Println(filename, "mode is", mode.String()[1:10])
}
