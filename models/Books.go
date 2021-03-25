package models

import "gorm.io/gorm"




type Book struct {
	gorm.Model
	Title  string `json:"Title"`
	Author string `json:"Author"`
	Rating int    `json:"Rating"`
}

type CreateBook struct {
	Title  string `json:"Title"`
	Author string `json:"Author"`
	Rating int    `json:"Rating"`
}