package main

import "C"
import "fmt"

//#include <stdio.h>
//void call(){
// printf("Calling C code!\n");
//}
func main() {
	fmt.Println("A go statment!!")
	C.call()
	fmt.Println("Anothger Go statment!")
}
