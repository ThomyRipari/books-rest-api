package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"math/rand"
	"strconv"
)

type Book struct {
	ID string `json: "id"`
	Isbn string `json: "isbn"`
	Title string `json: "title"`
	Author *Author `json: "author"`
}

type Author struct {
	Firstname string `json: "firstname"`
	Lastname string `json: "lastname"`
}

var books [] Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for i := 0; i < len(books); i++ {
		if books[i].ID == params["id"] {
			json.NewEncoder(w).Encode(books[i])
			return
		}
	}

	http.Error(w, "No se ha encontrado el libro", http.StatusNotFound)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book Book

	json.NewDecoder(r.Body).Decode(&book)

	book.ID = strconv.Itoa(rand.Intn(100000000))

	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for i := 0; i < len(books); i++ {
		if books[i].ID == params["id"] {
			books = append(books[:i], books[i + 1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(books)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for i := 0; i < len(books); i++ {
		if books[i].ID == params["id"] {
			books = append(books[:i], books[i + 1:]...)

			var book Book
			json.NewDecoder(r.Body).Decode(&book)

			book.ID = params["id"]
			books = append(books, book)

			json.NewEncoder(w).Encode(book)
			return

		}
	}
}

func main() {
	r := mux.NewRouter()

	books = append(books, Book{ID: "1", Isbn: "18921", Title: "Book One", Author: &Author{Firstname: "Author", Lastname: "One"}})
	books = append(books, Book{ID: "2", Isbn: "11562", Title: "Book Two", Author: &Author{Firstname: "Author", Lastname: "Two"}})

	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	http.ListenAndServe(":8000", r)

}