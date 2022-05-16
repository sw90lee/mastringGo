package main

import (
	"fmt"
	s "strings"
)

var f = fmt.Printf

func main() {
	upper := s.ToUpper("Hello there!")

	f("To Upper: %s\n", upper)
	f("To Lower: %s\n", s.ToLower("Hello THERE!!"))
	f("%s\n", s.Title("tHis wiLL be A title!"))
	f("EqualFold: %v\n", s.EqualFold("Mihalis", "MIHAlis"))
	f("EqualFold: %v\n", s.EqualFold("Mihalis", "MIHAli"))
}
