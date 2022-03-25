package bookrepo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	. "github.com/hkaya15/PicusSecurity/Week_4_Homework/book/model"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (b *BookRepository) GetBooksWithAuthor() ([]Book, error) {
	var result []Book

	if err := b.db.Preload("Authors").Find(&result).Error; err != nil {
		fmt.Println(err)
	}
	for _, v := range result {
		fmt.Println(v.ToString())
	}
	return result, nil
}

// BuyByID returns the book that buy by book ıd & book count if it causes negative value (uint)
func (b *BookRepository) BuyByID(id int, count uint) (Book, error) {
	var book Book
	b.db.Where(Book{BookID: id}).Find(&book)
	book.StockCount = book.StockCount - count
	err := b.db.Save(&book).Error
	if err != nil {
		log.Fatalln(err)
		return book, err
	}
	return book, nil
}

// DeleteByID returns the book that deleted by book id. It just update Deleted_At on DB (soft delete)
func (b *BookRepository) DeleteByID(id int) (Book, error) {
	var book Book
	result := b.db.Where(Book{BookID: id}).Find(&book).Model(&book).Update(os.Getenv("DELETED_AT"), time.Now())
	if result.Error != nil {
		log.Fatalln(result.Error.Error())
	}
	err := book.BeforeDelete(b.db)
	if err != nil {
		log.Fatalln(err)
		return book, err
	}
	errAfter := book.AfterDelete(b.db)
	if errAfter != nil {
		log.Fatalln(err)
		return book, err
	}
	return book, nil
}

// FindAll returns the book list
func (b *BookRepository) FindAll() []Book {
	var books []Book
	b.db.Find(&books)

	return books
}

// SearchName returns the book list by book name regarding of the contains & non-case sensitive
func (b *BookRepository) SearchByName(name string) []Book {
	var books []Book
	b.db.Where("name ILIKE ? ", "%"+name+"%").Find(&books)
	return books
}

func (b *BookRepository) Migrations() {
	b.db.AutoMigrate(&Book{})
}

func (b *BookRepository) InsertData() error {
	bookList, err := getAllBooksFromJSON()
	if err != nil {
		log.Fatalln(err)
		return err
	}

	for _, v := range bookList.Books {
		b.db.Where(Book{BookID: v.BookID}).Attrs(Book{BookID: v.BookID, BookName: v.BookName, Page: v.Page, StockCount: v.StockCount, Price: v.Price, StockID: v.StockID, ISBN: v.ISBN, AuthorID: v.AuthorID}).FirstOrCreate(&v)
	}
	return nil
}

// getAllBooksFromJSON returns the book list that readed by json
func getAllBooksFromJSON() (*Books, error) {

	var books Books
	jsonFile, err := os.Open(os.Getenv("BOOK_JSON"))
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	byteVal, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteVal, &books)

	defer jsonFile.Close()

	return &books, nil
}
