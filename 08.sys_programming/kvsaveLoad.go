package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

var DATAFILE = "./tmp/dataFile.gob"

func save() error {
	fmt.Println("Saving", DATAFILE)
	// 파일 삭제
	err := os.Remove(DATAFILE)
	if err != nil {
		fmt.Println(err)
	}

	saveTo, err := os.Create(DATAFILE)
	if err != nil {
		fmt.Println("Cannot create", DATAFILE)
		return err
	}
	defer saveTo.Close()

	encoder := gob.NewEncoder(saveTo)
	err = encoder.Encode(DATA)
	if err != nil {
		fmt.Println("Cannot save to", DATAFILE)
		return err
	}
	return nil
}

func load() error {
	fmt.Println("Loading", DATAFILE)
	// 파일을 읽음
	loadFrom, err := os.Open(DATAFILE)
	defer loadFrom.Close()

	if err != nil {
		fmt.Println("Empty Key/value store!")
		return err
	}

	decoder := gob.NewDecoder(loadFrom)
	decoder.Decode(&DATA)
	return nil
}
