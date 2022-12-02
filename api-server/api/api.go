package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/tapojit047/learning-go/api-server/db"
	"github.com/tapojit047/learning-go/api-server/handler"
	"log"
	"net/http"
)

func HandleRequests() {
	db.InitializeBooks()
	db.InitializeCred()
	r := chi.NewRouter()

	// Protected routes
	r.Group(func(r chi.Router) {
		// Seek, verify and validate JWT tokens
		r.Use(jwtauth.Verifier(db.TokenAuth))

		// Handle valid / invalid tokens. In this example, we use
		// the provided authenticator middleware, but you can write your
		// own very easily, look at the Authenticator method in jwtauth.go
		// and tweak it, its not scary.
		r.Use(jwtauth.Authenticator)

		//r.Post("/books", handler.AddBook())
		r.Put("/books/{id}", handler.UpdateBook())
		r.Delete("/books/{id}", handler.DeleteBook())
	})
	// Public routes
	r.Group(func(r chi.Router) {
		r.Post("/login", handler.Login)
		log.Println("In HandleRequest")
		r.Post("/books", handler.AddBook())
		r.Get("/books", handler.GetBooks())
		r.Get("/books/{id}", handler.GetBookById())
	})
	http.ListenAndServe(":8000", r)
}
