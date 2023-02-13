package questions

import (
	"encoding/csv"
	"io"
	"os"
)

func DumpFile(path string, questions []*Question) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return Dump(file, questions)
}

func Dump(writer io.Writer, questions []*Question) error {
	csvWriter := csv.NewWriter(writer)
	defer csvWriter.Flush()
	for _, question := range questions {
		dumped := dumpQuestion(question)
		err := csvWriter.Write(dumped)
		if err != nil {
			return err
		}
	}
	return nil
}

func dumpQuestion(question *Question) []string {
	return []string{question.Text, question.Answer}
}
