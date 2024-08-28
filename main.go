package main

import (
	"fmt"
	"github.com/fnoon/astrolib/num"
	"github.com/fnoon/astrolib/roman_date"
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
	date := roman_date.Date{
		Year:  2024,
		Month: roman_date.Aug,
		Day:   26,
	}

	fmt.Println(date)
	fmt.Println(num.Ceil(4.5))
	return nil
}
