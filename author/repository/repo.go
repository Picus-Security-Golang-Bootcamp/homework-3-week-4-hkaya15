package authorrepo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	. "github.com/hkaya15/PicusSecurity/Week_4_Homework/author/model"
	"gorm.io/gorm"
)

type AuthorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{db: db}
}

type AuthorsWithBook struct {
	gorm.Model
	ID int    `gorm:"foreignKey:ID"`
	Name   string `gorm:"foreignKey:name"`
	Page string `gorm:"foreignKey:page"`
	BookName string `gorm:"foreignKey:bookname"`
}

func (a *AuthorRepository) GetAuthorsWithBook() ([]AuthorsWithBook, error) {
	var author []AuthorsWithBook

	x := a.db.Table("authors").Select("authors.id,authors.name,books.page,books.book_name").Joins("Inner join books on books.author_id = authors.id")
	x.Find(&author)
	for i := 0; i < len(author); i++ {
		fmt.Println(author[i])
	}
	return author, nil
}
func (a *AuthorRepository) SearchByName(name string) []Author {
	var authors []Author
	fmt.Println(name)
	a.db.Where("Name ILIKE ? ", "%"+name+"%").Find(&authors)
	return authors
}

func (a *AuthorRepository) FindAll() []Author {
	var authors []Author
	a.db.Find(&authors)

	return authors
}

func (c *AuthorRepository) Migrations() {
	c.db.AutoMigrate(&Author{})
}

func (c *AuthorRepository) InsertData() error {
	authorsList, err := getAllAuthorsFromJSON()
	if err != nil {
		log.Fatalln(err)
		return err
	}

	for _, v := range authorsList.Authors {
		//fmt.Println(&v.Authorid)
		c.db.Where(Author{ID: v.ID}).Attrs(Author{ID: v.ID, Name: v.Name}).FirstOrCreate(&v)
	}
	return nil
}

func getAllAuthorsFromJSON() (*Authors, error) {

	var authors Authors
	jsonFile, err := os.Open(os.Getenv("AUTHOR_JSON"))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	byteVal, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteVal, &authors)

	defer jsonFile.Close()

	return &authors, nil
}
