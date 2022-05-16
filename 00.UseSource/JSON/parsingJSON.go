package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	argument := os.Args
	if len(argument) == 1 {
		fmt.Println("Please provide a filename!!")
		return
	}

	filename := argument[1]

	fileData, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	var parseData map[string]interface{}
	json.Unmarshal([]byte(fileData), &parseData)

	for key, value := range parseData {
		fmt.Println("key:", key, "value:", value)
	}
}
