package handler

import (
	"encoding/json"
	"github.com/tapojit047/learning-go/api-server/db"
	"github.com/tapojit047/learning-go/api-server/model"
	"log"
	"net/http"
)

func GetBooks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(db.Books)
	}
}

func GetBookById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := parseURL(r.URL.Path)
		for _, v := range db.Books {
			if v.ID == id {
				json.NewEncoder(w).Encode(v)
				return
			}
		}
		json.NewEncoder(w).Encode("Book Not Found")
	}
}

func AddBook() http.HandlerFunc {
	log.Println("In Addbook")
	return func(w http.ResponseWriter, r *http.Request) {
		request := map[string]string{}
		json.NewDecoder(r.Body).Decode(&request)

		newBook := model.Book{request["id"], request["title"], &model.Author{request["firstname"], request["lastname"]}}
		db.Books = append(db.Books, newBook)
	}
}

func UpdateBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := parseURL(r.URL.Path)
		request := map[string]string{}
		json.NewDecoder(r.Body).Decode(&request)
		newBook := model.Book{request["id"], request["title"], &model.Author{request["firstname"], request["lastname"]}}

		for i, bookVal := range db.Books {
			if bookVal.ID == id {
				db.Books[i] = newBook
				break
			}
		}
		json.NewEncoder(w).Encode(&newBook)
	}
}

func DeleteBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := parseURL(r.URL.Path)
		for i, bookVal := range db.Books {
			if bookVal.ID == id {
				db.Books[i] = db.Books[len(db.Books)-1]
				db.Books = db.Books[:len(db.Books)-1]
				break
			}
		}
	}
}
