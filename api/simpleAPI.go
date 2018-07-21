package api

import (
	"encoding/json"
	"log"
)

// Book is a data structures holding information about author, title, and ISBN of a book
type Book struct {
	Title  string `json:"Title"`
	Author string `json:"Author"`
	ISBN   string `json:"ISBN"`
}

// ToJSON can be used to marshal Book structs to JSON
func (b Book) ToJSON() []byte {
	json, err := json.Marshal(b)
	if err != nil {
		log.Fatalf("error marshalling Book: %v\n", err)
	}
	return json
}

// FromJSON can be used to unmarshal Book structs from JSON
func FromJSON(b []byte) Book {
	book := new(Book)
	err := json.Unmarshal(b, book)
	if err != nil {
		log.Fatalf("error unmarshalling Book: %v\n", err)
	}
	return *book
}
