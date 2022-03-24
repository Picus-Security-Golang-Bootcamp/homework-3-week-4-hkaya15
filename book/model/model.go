package bookmodel

import (
	"fmt"

	"gorm.io/gorm"
)

type Books struct {
	Books []Book `json:"Books"`
}

// Create Book type
type Book struct {
	gorm.Model
	BookID     int     `json:"BookID" gorm:"not_null"`
	Name       string  `json:"Name"`
	Page       int     `json:"Page"`
	StockCount int     `json:"StockCount"`
	Price      float64 `json:"Price" gorm:"type:numeric(10,2)"`
	StockID    int     `json:"StockID"`
	ISBN       int     `json:"ISBN"`
	AuthorID   int     `json:"AuthorID" gorm:"foreignKey:ID"`
}

func (Book) TableName() string {
	return "Book"
}

func (b *Book) ToString() string {
	return fmt.Sprintf("ID : %d, Name : %s, Page : %v, StockCount : %v,Price : %v, StockID : %v, ISBN : %v, AuthorID %v, CreatedAt : %s", b.BookID, b.Name, b.Page, b.StockCount, b.Price, b.StockID, b.ISBN, b.AuthorID, b.CreatedAt.Format("2006-01-02 15:04:05"))
}

func (b *Book) BeforeDelete(tx *gorm.DB) (err error) {
	fmt.Printf("Book (%s) deleting...", b.Name)
	return nil
}
