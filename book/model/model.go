package bookmodel

import . "github.com/hkaya15/PicusSecurity/Week_4_Homework/author/model"

type Books struct {
	Books []Book `json:"Books"`
}

// Create Book type
type Book struct {
	BookID     int     `json:"BookID"`
	Name       string  `json:"Name"`
	Page       int     `json:"Page"`
	StockCount int     `json:"StockCount"`
	Price      float64 `json:"Price"`
	StockID    int     `json:"StockID"`
	ISBN       int     `json:"ISBN"`
	Author     Author  `json:"Author"`
}
