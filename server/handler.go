package server

import (
	"adventurebook/adventurebook"
	"errors"
	"html/template"
	"net/http"
	"strings"
)

var storyTemplate = template.Must(template.ParseFiles("server/template.html"))

type Handler struct {
	books map[string]*adventurebook.AdventureBook
}

func NewHandler(books map[string]*adventurebook.AdventureBook) *Handler {
	return &Handler{
		books: books,
	}
}

func (h *Handler) getArc(r *http.Request) (*adventurebook.StoryArc, error) {
	bookName := strings.Trim(r.URL.Path, "/")
	book, ok := h.books[bookName]
	if !ok {
		return nil, errors.New("book not found")
	}
	arcName := r.URL.Query().Get("arc")
	if arcName == "" {
		arcName = adventurebook.DEFALUT_ARC_NAME
	}
	return book.GetArc(arcName)
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	arc, err := h.getArc(r)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	storyTemplate.Execute(w, arc)
}
