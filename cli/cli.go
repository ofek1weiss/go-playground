package cli

import (
	"adventurebook/adventurebook"
	"fmt"
	"strconv"
)

type Cli struct {
	book *adventurebook.AdventureBook
	arc  *adventurebook.StoryArc
}

func New(book *adventurebook.AdventureBook) *Cli {
	return &Cli{
		book: book,
		arc:  book.GetFirstArc(),
	}
}

func (c *Cli) didFinish() bool {
	return len(c.arc.Options) == 0
}

func (c *Cli) printArc() {
	fmt.Println(c.arc.Title)
	fmt.Println()
	for _, paragraph := range c.arc.Story {
		fmt.Println(paragraph)
	}
	fmt.Println()
	for i, option := range c.arc.Options {
		fmt.Printf("%d - %s\n", i+1, option.Text)
	}
}

func (c *Cli) getOption() *adventurebook.Option {
	minOption := 1
	maxOption := len(c.arc.Options)
	var input string
	for {
		fmt.Println("What will you do?")
		fmt.Scanln(&input)
		optionIndex, err := strconv.Atoi(input)
		if err != nil || optionIndex < minOption || optionIndex > maxOption {
			fmt.Println("Invalid option, try again...")
		} else {
			return c.arc.Options[optionIndex-1]
		}
	}
}

func (c *Cli) switchToOption(chosenOption *adventurebook.Option) {
	arc, err := c.book.GetArc(chosenOption.Arc)
	if err != nil {
		panic(err)
	}
	c.arc = arc
}

func (c *Cli) Run() {
	for !c.didFinish() {
		c.printArc()
		chosenOption := c.getOption()
		c.switchToOption(chosenOption)
	}
	c.printArc()
	fmt.Println("The end")
}
