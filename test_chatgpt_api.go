package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Json struct {
	Text string `json:"text"`
}

func read(filename string) (string, error) {
	file, err := os.Open(fmt.Sprintf("%s.json", filename))
	if err != nil {
		return "", err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return "", err
	}

	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return "", err
	}

	return string(bs), nil
}

func readJson(filename string) (Json, error) {
	var jsn Json
	file, err := os.Open(filename)
	if err != nil {
		return jsn, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&jsn)
	if err != nil {
		return jsn, err
	}

	return jsn, nil
}

func main() {
	jsn, err := read("l1")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(jsn)
}
