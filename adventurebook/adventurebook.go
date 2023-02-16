package adventurebook

import (
	"encoding/json"
	"errors"
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

type AdventureBook struct {
	arcs map[string]*StoryArc
}

func New(arcs map[string]*StoryArc) (*AdventureBook, error) {
	_, hasDefaultArc := arcs[DEFALUT_ARC_NAME]
	if !hasDefaultArc {
		return nil, errors.New("missing default arc")
	}
	return &AdventureBook{arcs: arcs}, nil
}

func LoadFile(path string) (*AdventureBook, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var arcs map[string]*StoryArc
	err = json.Unmarshal(data, &arcs)
	if err != nil {
		return nil, err
	}
	return New(arcs)
}

func (ab *AdventureBook) GetArc(arcName string) (*StoryArc, error) {
	arc, ok := ab.arcs[arcName]
	if !ok {
		return nil, errors.New("arc not found")
	}
	return arc, nil
}

func (ab *AdventureBook) GetFirstArc() *StoryArc {
	return ab.arcs[DEFALUT_ARC_NAME]
}
