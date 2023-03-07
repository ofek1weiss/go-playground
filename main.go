package main

import (
	"fmt"
	"os"
	"task/cli"
)

func main() {
	if err := cli.App.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
