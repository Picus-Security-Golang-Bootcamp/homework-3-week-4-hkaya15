package authormodel

import "gorm.io/gorm"

type Authors struct {
	Authors []Author `json:"Authors"`
}

// Create Author type
type Author struct {
	gorm.Model
	ID   int    `json:"AuthorID"`
	Name string `json:"Name"`
}

