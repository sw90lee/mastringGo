package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// JSON 데이터 저장

type record struct {
	Name    string
	Surname string
	Tel     []telephone
}

type telephone struct {
	Mobile bool
	Number string
}

func saveToJSON(filename *os.File, key interface{}) {
	encodeJSON := json.NewEncoder(filename)
	err := encodeJSON.Encode(key)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	myRecord := record{
		Name:    "Sungwoo",
		Surname: "Lee",
		Tel: []telephone{
			telephone{Mobile: true, Number: "1234-567"},
			telephone{Mobile: true, Number: "1234-abcd"},
			telephone{Mobile: false, Number: "1234-567"},
		},
	}

	saveToJSON(os.Stdout, myRecord)
}
