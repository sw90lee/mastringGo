package main

import (
	"flag"
	"fmt"
	"strings"
)

type Value interface {
	String() string
	Set(string) error
}

type NamesFlag struct {
	Names []string
}

func (s *NamesFlag) GetNames() []string {
	return s.Names
}

func (s *NamesFlag) String() string {
	return fmt.Sprint(s.Names)
}

func (s *NamesFlag) Set(v string) error {
	if len(s.Names) > 0 {
		return fmt.Errorf("Cannot use names flag more the once!")
	}

	names := strings.Split(v, ",")
	for _, item := range names {
		s.Names = append(s.Names, item)
	}
	return nil
}

func main() {
	var manyNames NamesFlag

	minusK := flag.Int("k", 0, "An int")
	minusO := flag.String("o", "Sungwoo", "The name")
	flag.Var(&manyNames, "names", "Comma-separted list")
	flag.Parse()

	fmt.Println("-k", *minusK)
	fmt.Println("-o", *minusO)

	for i, item := range manyNames.GetNames() {
		fmt.Println(i, item)
	}

	fmt.Println("Remaining command-line arguments:")
	for index, val := range flag.Args() {
		fmt.Println(index, ":", val)
	}
}
