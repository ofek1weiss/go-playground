package main

import "flag"

type Args struct {
	QuestionsFile string
	Timeout       int
	Shuffle       bool
}

func ParseArgs() *Args {
	var args Args
	flag.StringVar(&args.QuestionsFile, "questions", "questions.csv", "")
	flag.IntVar(&args.Timeout, "timeout", 30, "")
	flag.BoolVar(&args.Shuffle, "shuffle", false, "")
	flag.Parse()
	return &args
}
