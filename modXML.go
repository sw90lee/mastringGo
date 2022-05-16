package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

func main() {
	type Address struct {
		City, Country string
	}

	type Employee struct {
		XMLName   xml.Name `xml: "employee`
		Id        int      `xml: "id,attr"`
		FirstName string   `xml :"name>first"` //first 태그를 Name태그 안에 넣기
		LastName  string   `xml :"name>last"`  // last 태그를 Name태그 안에 넣기
		Initials  string   `xml :"name>initials"`
		Height    float32  `xml:"height,omitempty"` //omitempty 필드가 공백일 경우 출력 생략
		Address
		Comment string `xml:",comment"` //,comment 주석에 관한 필드 <!-- ㅁㄴㅇㄻㅇㄹ --> 으로 출력
	}

	r := &Employee{Id: 7, FirstName: "Lee", LastName: "Sungwoo", Initials: "LSW"}
	r.Comment = "Technical Writer + Devops"
	r.Address = Address{"SomeWhere 12", "12312, Seoul"}

	output, err := xml.MarshalIndent(r, "  ", "    ")
	if err != nil {
		fmt.Println("Error:", err)
	}
	output = []byte(xml.Header + string(output))
	os.Stdout.Write(output)
	os.Stdout.Write([]byte("\n"))
}
