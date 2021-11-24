package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Book Structs (model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

//Author Struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init books var as a slice book struct
var books []Book

//Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)

}

//get single book
func getBook(w http.ResponseWriter, r *http.Request) {

}

//create single book
func createBook(w http.ResponseWriter, r *http.Request) {

}

//update single book
func updateBook(w http.ResponseWriter, r *http.Request) {

}

//delete single book
func deleteBook(w http.ResponseWriter, r *http.Request) {

}

func main() {
	//Init ruoter
	r := mux.NewRouter()

	//Mock Data
	books = append(books, Book{ID: "1", Isbn: "445258", Title: "Dracula", Author: &Author{Firstname: "Bram", Lastname: "Stoker"}})
	books = append(books, Book{ID: "2", Isbn: "782155", Title: "Return of the King", Author: &Author{Firstname: "J.R.R", Lastname: "Tolkien"}})
	books = append(books, Book{ID: "3", Isbn: "332541", Title: "Foundation", Author: &Author{Firstname: "Isaac", Lastname: "Isamov"}})
	books = append(books, Book{ID: "4", Isbn: "685218", Title: "SilverThorn", Author: &Author{Firstname: "Raymond", Lastname: "Fiest"}})

	// route Handlers / Endpoints

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("Delete")

	log.Fatal(http.ListenAndServe(":8000", r))
}
