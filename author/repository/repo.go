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

// GetAuthorsWithBooks return DTO combined with book names
func (a *AuthorRepository) GetAuthorsWithBooks() ([]AuthorsWithBook, error) {
	var author []AuthorsWithBook
	res := a.db.Table("authors").Select("authors.id,authors.name,books.page,books.book_name").Joins("Inner join books on books.author_id = authors.id")
	res.Find(&author)
	for i := 0; i < len(author); i++ {
		fmt.Println(author[i])
	}
	return author, nil
}

// SearchByName returns Authors list by name
func (a *AuthorRepository) SearchByName(name string) []Author {
	var authors []Author
	fmt.Println(name)
	a.db.Where("Name ILIKE ? ", "%"+name+"%").Find(&authors)
	return authors
}

// FindAll returns all author list
func (a *AuthorRepository) FindAll() []Author {
	var authors []Author
	a.db.Find(&authors)

	return authors
}

// Migrations helps the auto-migrate
func (c *AuthorRepository) Migrations() {
	c.db.AutoMigrate(&Author{})
}

// InsertData helps the insert data
func (c *AuthorRepository) InsertData() error {
	authorsList, err := getAllAuthorsFromJSON()
	if err != nil {
		log.Fatalln(err)
		return err
	}

	for _, v := range authorsList.Authors {
		c.db.Where(Author{ID: v.ID}).Attrs(Author{ID: v.ID, Name: v.Name}).FirstOrCreate(&v)
	}
	return nil
}

// getAllAuthorsFromJSON convert JSON to Author slice
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
