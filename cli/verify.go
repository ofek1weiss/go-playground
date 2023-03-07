package cli

import (
	"errors"

	"github.com/urfave/cli/v2"
)

func verifyArgs(args cli.Args) error {
	if !args.Present() {
		return errors.New("no args supplied")
	}
	return nil
}

func verifyNoArgs(args cli.Args) error {
	if args.Present() {
		return errors.New("args not supported")
	}
	return nil
}
