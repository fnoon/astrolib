package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	os.Exit(run(os.Args, os.Stdin, os.Stderr, os.Stdout))
}

func run(args []string, stdin io.Reader, stderr io.Writer, stdout io.Writer) int {
	// ...
	fmt.Printf("Hello: %v\n", args)
	return 0
}
