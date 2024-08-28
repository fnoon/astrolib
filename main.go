package main

import (
	"fmt"
	"github.com/fnoon/astrolib/date"
	"io"
	"os"
)

func main() {
	err := run(os.Args, os.Stdin, os.Stderr, os.Stdout)
	var status int
	if err != nil {
		status = 1
	} else {
		status = 0
	}
	os.Exit(status)
}

func run(args []string, stdin io.Reader, stderr io.Writer, stdout io.Writer) error {
	// ...
	fmt.Printf("Hello: %v\n", args)
	date := date.Date{
		Year:  2024,
		Month: date.Aug,
		Day:   26,
	}

	fmt.Print(date)
	return nil
}
