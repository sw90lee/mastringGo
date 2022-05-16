package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
)

type Record struct {
	Name    string
	Surname string
	Tel     []Telephone
}

type Telephone struct {
	Mobile bool
	Number string
}

// JSON 파일 읽기
func loadFromJSON(filename string, key interface{}) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}

	decodeJSON := json.NewDecoder(in)
	err = decodeJSON.Decode(key)
	if err != nil {
		return err
	}
	in.Close()
	return nil
}

func main() {
	arguement := os.Args
	if len(arguement) == 1 {
		fmt.Println("Please provide a filename!")
		return
	}

	filename := arguement[1]

	var myRecord Record
	err := loadFromJSON(filename, &myRecord)
	if err == nil {
		fmt.Println("JSON:", myRecord)
	} else {
		fmt.Println(err)
	}

	myRecord.Name = "Sungwoo"

	// XML 변환
	xmlData, _ := xml.MarshalIndent(myRecord, "", "    ")
	// XML Header 추가
	xmlData = []byte(xml.Header + string(xmlData))
	fmt.Println("\nxmlData:", string(xmlData))

	data := &Record{}
	err = xml.Unmarshal(xmlData, data)
	if nil != err {
		fmt.Println("Unmarshalling from XML", err)
		return
	}

	result, err := json.Marshal(data)
	if nil != err {
		fmt.Println("Error marshaling to Json", err)
		return
	}

	_ = json.Unmarshal([]byte(result), &myRecord)
	fmt.Println("\nJSON:", myRecord)
}
