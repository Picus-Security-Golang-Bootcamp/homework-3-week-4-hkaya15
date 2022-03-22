package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	. "github.com/hkaya15/PicusSecurity/Week_4_Homework/book/model"
)

func main() {
	x, e := GetAllBooks()
	fmt.Println(x, e)
}
func GetAllBooks() (*Books, error) {

	var books Books
	jsonFile, err := os.Open("source/books.json")
	if err != nil {
		return nil, err
	}

	byteVal, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteVal, &books)

	defer jsonFile.Close()
	if err != nil {
		fmt.Println("Error:", err)
	}

	return &books, nil
}
