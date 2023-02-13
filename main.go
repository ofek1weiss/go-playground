package main

import (
	"playground/quiz"
	"playground/quiz/questions"
)

func main() {
	args := ParseArgs()
	questions, err := questions.LoadFile(args.QuestionsFile)
	if err != nil {
		panic(err)
	}
	quizGame := quiz.NewQuiz(questions, args.Timeout)
	if args.Shuffle {
		quizGame.Shuffle()
	}
	err = quizGame.Run()
	if err != nil {
		panic(err)
	}
}
