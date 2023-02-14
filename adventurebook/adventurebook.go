package adventurebook

import (
	"encoding/json"
	"os"
)

const DEFALUT_ARC_NAME = "intro"

type Option struct {
	Text string `json:text`
	Arc  string `json:arc`
}

type StoryArc struct {
	Title   string    `json:title`
	Story   []string  `json:story`
	Options []*Option `json:options`
}

type AdventureBook map[string]*StoryArc

func LoadFile(path string) (*AdventureBook, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var book AdventureBook
	err = json.Unmarshal(data, &book)
	if err != nil {
		return nil, err
	}
	return &book, nil
}
