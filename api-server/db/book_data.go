package db

import (
	"github.com/tapojit047/learning-go/api-server/model"
)

var Books []model.Book

func InitializeBooks() {
	Books = []model.Book{
		{ID: "1", Title: "Harry Potter: Chamber of Secrets", Author: &model.Author{FirstName: "J. K.", LastName: "Rowling"}},
		{ID: "2", Title: "The Lord of the Rings", Author: &model.Author{FirstName: "J. R. R.", LastName: "Tolkien"}},
		{"3", "A Song of Ice and Fire", &model.Author{FirstName: "George R. R.", LastName: "Martin"}},
		{"4", "The Alchemist", &model.Author{FirstName: "George R. R.", LastName: "Martin"}},
	}
}
