package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"text/template"
)

type Entry struct {
	Number int
	Double int
	Squere int
}

var DATA []Entry
var tFile string

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Host: %s Path: %s\n", r.Host, r.URL.Path)
	myT := template.Must(template.ParseGlob(tFile))
	myT.ExecuteTemplate(w, tFile, DATA)
}

func main() {
	argument := os.Args
	if len(argument) != 3 {
		fmt.Println("Need Database file + Template File!")
		return
	}

	database := argument[1]
	tFile = argument[2]

	// 데이터베이스 연결
	db, err := sql.Open("sqlite3", database)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Emptying database table!")
	// 데이터베이스 실행
	_, err = db.Exec("DELETE FROM data")
	if err != nil {
		fmt.Println(nil)
		return
	}

	fmt.Println("Pipulating", database)
	stmt, _ := db.Prepare("INSTER INTO data(number, double, squere) value (?,?,?)")
	for i := 20; i < 50; i++ {
		_, _ = stmt.Exec(i, 2*i, i*i)
	}

	rows, err := db.Query("SELECT * FROM data")
	if err != nil {
		fmt.Println(nil)
		return
	}

	var n int
	var d int
	var s int
	for rows.Next() {
		err = rows.Scan(&n, &d, &s)
		temp := Entry{Number: n, Double: d, Squere: s}
		DATA = append(DATA, temp)
	}

	http.HandleFunc("/", myHandler)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
