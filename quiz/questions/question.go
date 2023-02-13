package questions

import "strings"

type Question struct {
	Text   string
	Answer string
}

func normalizeAnswer(str string) string {
	return strings.ToLower(strings.TrimSpace(str))
}

func NewQuestion(text string, answer string) *Question {
	return &Question{
		Text:   text,
		Answer: answer,
	}
}

func (q *Question) IsAnswerCorrect(answer string) bool {
	return normalizeAnswer(q.Answer) == normalizeAnswer(answer)
}
