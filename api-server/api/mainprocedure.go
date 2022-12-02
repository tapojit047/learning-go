package api

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"log"
	"net/http"
	"strings"
	"time"
)

type book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastName"`
}

type credentials struct {
	Username string `json:"username"'`
	Password string `json:"password"`
}

var credentialList map[string]string
var books []book
var tokenAuth *jwtauth.JWTAuth

func getBooks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(books)
	}
}

func getBookById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := parseURL(r.URL.Path)
		for _, v := range books {
			if v.ID == id {
				json.NewEncoder(w).Encode(v)
				return
			}
		}
		json.NewEncoder(w).Encode("Book Not Found")
	}
}

func addBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := map[string]string{}
		json.NewDecoder(r.Body).Decode(&request)

		newBook := book{request["id"], request["title"], &Author{request["firstname"], request["lastname"]}}
		books = append(books, newBook)
	}
}

func parseURL(url string) string {
	p := strings.Split(url, "/")
	return p[len(p)-1]
}

func updateBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := parseURL(r.URL.Path)
		request := map[string]string{}
		json.NewDecoder(r.Body).Decode(&request)
		newBook := book{request["id"], request["title"], &Author{request["firstname"], request["lastname"]}}

		for i, bookVal := range books {
			if bookVal.ID == id {
				books[i] = newBook
				break
			}
		}
		json.NewEncoder(w).Encode(&newBook)
	}
}

func deleteBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := parseURL(r.URL.Path)
		for i, bookVal := range books {
			if bookVal.ID == id {
				books[i] = books[len(books)-1]
				books = books[:len(books)-1]
				break
			}
		}
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	var info credentials

	json.NewDecoder(r.Body).Decode(&info)

	pass, ok := credentialList[info.Username]
	log.Println(info.Password, info.Username, pass, ok)
	if !ok || pass != info.Password {
		json.NewEncoder(w).Encode("Invalid Authorization")
		return
	}
	expiretime := time.Now().Add(15 * time.Minute)
	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"user_id": 123})

	http.SetCookie(w, &http.Cookie{
		Name:    "cookie",
		Value:   tokenString,
		Expires: expiretime,
	})
}

func initialize() {
	books = []book{
		{ID: "1", Title: "Harry Potter: Chamber of Secrets", Author: &Author{FirstName: "J. K.", LastName: "Rowling"}},
		{ID: "2", Title: "The Lord of the Rings", Author: &Author{FirstName: "J. R. R.", LastName: "Tolkien"}},
		{"3", "A Song of Ice and Fire", &Author{FirstName: "George R. R.", LastName: "Martin"}},
		{"4", "The Alchemist", &Author{FirstName: "George R. R.", LastName: "Martin"}},
	}
	credentialList = map[string]string{
		"tapojit047": "1234",
	}
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)
}

func handleRequests() {
	r := chi.NewRouter()

	// Protected routes
	r.Group(func(r chi.Router) {
		// Seek, verify and validate JWT tokens
		r.Use(jwtauth.Verifier(tokenAuth))

		// Handle valid / invalid tokens. In this example, we use
		// the provided authenticator middleware, but you can write your
		// own very easily, look at the Authenticator method in jwtauth.go
		// and tweak it, its not scary.
		r.Use(jwtauth.Authenticator)

		r.Post("/books", addBook())
		r.Put("/books/{id}", updateBook())
		r.Delete("/books/{id}", deleteBook())
	})
	// Public routes
	r.Group(func(r chi.Router) {
		r.Post("/login", Login)
		r.Get("/books", getBooks())
		r.Get("/books/{id}", getBookById())
	})
	http.ListenAndServe(":8000", r)
}

func MainProcedure() {
	initialize()
	handleRequests()
}
