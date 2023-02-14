package main

import (
	"adventurebook/adventurebook"
	"adventurebook/server"
	"log"
	"net/http"
)

func getBooks() map[string]*adventurebook.AdventureBook {
	story, err := adventurebook.LoadFile("stories/gopher.json")
	if err != nil {
		panic(err)
	}
	return map[string]*adventurebook.AdventureBook{
		"gopher": story,
	}
}

func main() {
	handler := server.NewHandler(getBooks())
	http.Handle("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
