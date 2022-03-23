package bookmodel

import "gorm.io/gorm"

type Books struct {
	Books []Book `json:"Books"`
}

// Create Book type
type Book struct {
	gorm.Model
	BookID     int     `json:"BookID" gorm:"not_null; column: BookID"`
	Name       string  `json:"Name" gorm:"column: Name"`
	Page       int     `json:"Page" gorm:"column: Page"`
	StockCount int     `json:"StockCount" gorm:"column: StockCount"`
	Price      float64 `json:"Price" gorm:"type:numeric(10,2); column: Price"`
	StockID    int     `json:"StockID" gorm:"column: StockID"`
	ISBN       int     `json:"ISBN" gorm:"column: ISBN"`
	AuthorID   int     `json:"AuthorID" gorm:"foreignKey:ID; column: AuthorID"`
}

func (Book) TableName() string {
	return "Books"
}
