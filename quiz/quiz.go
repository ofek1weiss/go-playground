package quiz

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"playground/quiz/questions"
	"time"
)

type Quiz struct {
	Questions            []*questions.Question
	Timeout              int
	score                int
	currentQuestionIndex int
}

func NewQuiz(questions []*questions.Question, timeout int) *Quiz {
	return &Quiz{
		Questions:            questions,
		Timeout:              timeout,
		score:                0,
		currentQuestionIndex: 0,
	}
}

func (q *Quiz) Shuffle() {
	swap := func(i, j int) { q.Questions[i], q.Questions[j] = q.Questions[j], q.Questions[i] }
	rand.Shuffle(len(q.Questions), swap)
}

func (q *Quiz) Run() error {
	if len(q.Questions) == 0 {
		return errors.New("Quiz without questions cannot run")
	}
	fmt.Println("Welcome to the Quiz! lets begin...")
	done := make(chan bool)
	go func() {
		q.runAllQuestions()
		done <- true
	}()
	select {
	case <-done:
		fmt.Println("Thats all :)")
	case <-time.After(time.Duration(q.Timeout) * time.Second):
		fmt.Println("Time Out!")
	}
	fmt.Printf("Quiz Over, your score is %d/%d", q.score, len(q.Questions))
	return nil
}

func (q *Quiz) runAllQuestions() {
	question := q.getNextQuestion()
	for question != nil {
		q.runQuestion(question)
		question = q.getNextQuestion()
		fmt.Println()
	}
}

func (q *Quiz) getNextQuestion() *questions.Question {
	if q.currentQuestionIndex >= len(q.Questions) {
		return nil
	}
	question := q.Questions[q.currentQuestionIndex]
	q.currentQuestionIndex++
	return question
}

func (q *Quiz) runQuestion(question *questions.Question) {
	fmt.Println("Your next question is:", question.Text)
	fmt.Print("Your answer: ")
	answer, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	q.handleAnswer(question, answer)
}

func (q *Quiz) handleAnswer(question *questions.Question, answer string) {
	isCorrect := question.IsAnswerCorrect(answer)
	if isCorrect {
		fmt.Println("Correct!")
		q.score++
	} else {
		fmt.Println("Wrong!")
	}
}
