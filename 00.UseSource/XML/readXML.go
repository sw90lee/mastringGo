package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type record struct {
	Name    string
	Surname string
	Tel     []Telephone
}

type telephone struct {
	Mobile bool
	Number string
}

func loadFromXML(filename string, key interface{}) error {
	in, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return err
	}

	decodeXML := xml.NewDecoder(in)
	err = decodeXML.Decode(key)
	if err != nil {
		return err
	}
	in.Close()
	return nil
}

func main() {
	argument := os.Args
	if len(argument) == 1 {
		fmt.Println("Please provide a filename")
		return
	}

	filename := argument[1]

	var myRecord record
	err := loadFromXML(filename, &myRecord)

	if err == nil {
		fmt.Println("XML:", myRecord)
	} else {
		fmt.Println(err)
	}
}
