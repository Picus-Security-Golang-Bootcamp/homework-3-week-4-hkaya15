package bookrepo

import (
	"encoding/json"
	"fmt"
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

func (b *BookRepository) FindAll() []Book {
	var books []Book
	b.db.Find(&books)

	return books
}

func (b *BookRepository) SearchByName(name string) []Book {
	var books []Book
	//name = "%" + name + "%"
	b.db.Where("name ILIKE ? ", "%"+name+"%").Find(&books)
	fmt.Println(books)

	return books
}

func (b *BookRepository) Migrations() {
	b.db.AutoMigrate(&Book{})
}

func (b *BookRepository) InsertData() error {
	bookList, err := getAllBooksFromJSON()
	if err != nil {
		return err
	}

	for _, v := range bookList.Books {
		b.db.Where(Book{BookID: v.BookID}).Attrs(Book{BookID: v.BookID, Name: v.Name, Page: v.Page, StockCount: v.StockCount, Price: v.Price, StockID: v.StockID, ISBN: v.ISBN, AuthorID: v.AuthorID}).FirstOrCreate(&v)
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
