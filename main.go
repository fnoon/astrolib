package main

import (
	"fmt"
	"github.com/fnoon/astrolib/date"
	"io"
	"os"
)

func main() {
	os.Exit(run(os.Args, os.Stdin, os.Stderr, os.Stdout))
}

func run(args []string, stdin io.Reader, stderr io.Writer, stdout io.Writer) int {
	// ...
	fmt.Printf("Hello: %v\n", args)
	date := date.Date{
		Year:  2024,
		Month: date.Aug,
		Day:   26,
	}

	fmt.Print(date)
	return 0
}
