package questions

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func LoadFile(path string) ([]*Question, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return Load(file)
}

func Load(reader io.Reader) ([]*Question, error) {
	csvReader := csv.NewReader(reader)
	rawQuesions, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}
	questions := make([]*Question, len(rawQuesions))
	for i, rawQuestion := range rawQuesions {
		question, err := loadRecord(rawQuestion)
		if err != nil {
			return questions, err
		}
		questions[i] = question
	}
	return questions, nil
}

func loadRecord(record []string) (*Question, error) {
	if len(record) != 2 {
		return nil, fmt.Errorf("%v is not in the right length to be converted to a Question", record)
	}
	text, answer := record[0], record[1]
	return NewQuestion(text, answer), nil
}
