package main

import (
	"log"

	"github.com/hkaya15/PicusSecurity/Week_4_Homework/base/db"

	. "github.com/hkaya15/PicusSecurity/Week_4_Homework/author/repository"
	. "github.com/hkaya15/PicusSecurity/Week_4_Homework/book/repository"

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
	db, err := db.CreatePostgreSQL()
	if err != nil {
		log.Fatal("DB cannot init")
	}
	log.Println("Postgres connected: ", db)
	bookRepo := NewBookRepository(db)
	authorRepo := NewAuthorRepository(db)
	bookRepo.Migrations()
	authorRepo.Migrations()
	bookRepo.InsertData()
	authorRepo.InsertData()
}
