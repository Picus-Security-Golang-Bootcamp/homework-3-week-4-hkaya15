package authormodel

import "gorm.io/gorm"

type Authors struct {
	Authors []Author `json:"Authors"`
}

// Create Author type
type Author struct {
	gorm.Model
	ID   int    `json:"ID"`
	Name string `json:"Name"`
}

func (Author) TableName() string {
	return "Author"
}
