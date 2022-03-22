package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/hkaya15/PicusSecurity/Week_4_Homework/base/db"
	. "github.com/hkaya15/PicusSecurity/Week_4_Homework/book/model"
	"github.com/joho/godotenv"
)

func main() {
	//x, e := GetAllBooks()
	//fmt.Println(x, e)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//fmt.Println(os.Getenv("PICUS_DB_NAME"))
	d, err := db.CreatePostgreSQL()
	if err != nil {
		log.Fatal("DB cannot init")
	}
	log.Println("Postgres connected: ", d)
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
