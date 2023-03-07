package types

import "time"

type Task struct {
	Id             ID     `json:id`
	Text           string `json:text`
	CompletionTime *Time  `json:completionTime`
}

func NewTask(text string) *Task {
	return &Task{
		Id:             0,
		Text:           text,
		CompletionTime: nil,
	}
}

func (t *Task) Complete() {
	completionTime := Time(time.Now())
	t.CompletionTime = &completionTime
}

func (t *Task) IsComplete() bool {
	return t.CompletionTime != nil
}
