package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// JSON 데이터 읽기 예제

type Record struct {
	Name    string
	Surname string
	Tel     []Telephone
}

type Telephone struct {
	Mobile bool
	Number string
}

func loadFromJSON(filename string, key interface{}) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer in.Close()

	decodeJson := json.NewDecoder(in)

	if err = decodeJson.Decode(key); err != nil {
		return err
	}

	return nil
}

func main() {
	argument := os.Args

	if len(argument) == 1 {
		fmt.Println("Please provide a filename!")
		return
	}

	filename := argument[1]

	var myRecord Record
	err := loadFromJSON(filename, &myRecord)
	if err == nil {
		fmt.Println(myRecord)
	} else {
		fmt.Println(err)
	}
}
