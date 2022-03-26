package main

import (
	"log"

	. "github.com/hkaya15/PicusSecurity/Week_4_Homework/base/db"

	. "github.com/hkaya15/PicusSecurity/Week_4_Homework/author/repository"
	. "github.com/hkaya15/PicusSecurity/Week_4_Homework/book/repository"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}
	base := DBBase{DbType: &POSTGRES{}}
	db, err := base.DbType.Create()

	if err != nil {
		log.Fatalln("DB cannot init")
	}
	log.Println("DB connected: ", db)

	bookRepo := NewBookRepository(db)
	authorRepo := NewAuthorRepository(db)
	bookRepo.Migrations()
	authorRepo.Migrations()
	bookRepo.InsertData()
	authorRepo.InsertData()

}
