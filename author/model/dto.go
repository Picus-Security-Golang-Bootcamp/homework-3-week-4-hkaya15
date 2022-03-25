package authormodel

import "gorm.io/gorm"

type AuthorsWithBook struct {
	gorm.Model
	ID       int    `gorm:"foreignKey:ID"`
	Name     string `gorm:"foreignKey:name"`
	Page     string `gorm:"foreignKey:page"`
	BookName string `gorm:"foreignKey:bookname"`
}
