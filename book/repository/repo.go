package bookrepo

import (
	"encoding/json"
	"io/ioutil"
	"os"

	. "github.com/hkaya15/PicusSecurity/Week_4_Homework/book/model"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (c *BookRepository) Migrations() {
	c.db.AutoMigrate(&Book{})
}

func (c *BookRepository) InsertData() error {
	bookList, err := getAllBooksFromJSON()
	if err != nil {
		return err
	}

	for _, v := range bookList.Books {
		c.db.Where(Book{BookID: v.BookID}).Attrs(Book{BookID: v.BookID, Name: v.Name, Page: v.Page, StockCount: v.StockCount, Price: v.Price, StockID: v.StockID, ISBN: v.ISBN, AuthorID: v.AuthorID}).FirstOrCreate(&v)
	}
	return nil
}

func getAllBooksFromJSON() (*Books, error) {

	var books Books
	jsonFile, err := os.Open(os.Getenv("BOOK_JSON"))
	if err != nil {
		return nil, err
	}

	byteVal, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteVal, &books)

	defer jsonFile.Close()

	return &books, nil
}
