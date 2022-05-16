package main

import "C"
import "fmt"

//#include <stdio.h>
//void call(){
// printf("Calling C code!\n");
//}

func PrintMessage() {
	fmt.Println("A go Function!")
}

func Multiply(a,b int) int {
	return a*b
}

func main() {
}
