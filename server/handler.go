package server

import (
	"adventurebook/adventurebook"
	"html/template"
	"net/http"
	"strings"
)

var storyTemplate = template.Must(template.ParseFiles("server/template.html"))

type Handler struct {
	stories map[string]*adventurebook.AdventureBook
}

func NewHandler(stories map[string]*adventurebook.AdventureBook) *Handler {
	return &Handler{
		stories: stories,
	}
}

func (h *Handler) getArc(r *http.Request) (*adventurebook.StoryArc, bool) {
	storyName := strings.Trim(r.URL.Path, "/")
	story, ok := h.stories[storyName]
	if !ok {
		return nil, false
	}
	arcName := r.URL.Query().Get("arc")
	if arcName == "" {
		arcName = adventurebook.DEFALUT_ARC_NAME
	}
	arc, ok := (*story)[arcName]
	return arc, ok
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	arc, ok := h.getArc(r)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	storyTemplate.Execute(w, arc)
}
