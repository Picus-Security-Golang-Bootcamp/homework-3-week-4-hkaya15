package authorrepo

import (
	"encoding/json"
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
