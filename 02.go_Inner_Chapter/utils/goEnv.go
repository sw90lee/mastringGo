package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("You're using ", runtime.Compiler, " ")
	fmt.Println("on a", runtime.GOARCH, "machine")
	fmt.Println("using Go version", runtime.Version())
	fmt.Println("Number of CPUs:", runtime.NumCPU())
	fmt.Println("Number of Goroutines", runtime.NumGoroutine())
}
