package main

import "flag"

type Flags struct {
	Domain string
	Depth  int
}

func GetFlags() *Flags {
	var flags Flags
	flag.StringVar(&flags.Domain, "domain", "gobyexample.com", "")
	flag.IntVar(&flags.Depth, "depth", 10, "")
	flag.Parse()
	return &flags
}
