package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	lab2 "github.com/hrystynaa/lab2-go"
)

var (
	exprPtr = flag.String("e", "", "expression to evaluate")
	filePtr = flag.String("f", "", "file containing expression to evaluate")
	outPtr  = flag.String("o", "", "file to output result to ")
)

func main() {

	flag.Parse()

	var input io.Reader
	if *exprPtr != "" {
		input = strings.NewReader(*exprPtr)
	} else if *filePtr != "" {
		file, err := os.Open(*filePtr)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error opening file:", err)
			os.Exit(1)
		}
		defer file.Close()
		input = file
	} else {
		fmt.Fprintln(os.Stderr, "error: no input source specified")
		os.Exit(1)
	}

	var output io.Writer
	if *outPtr != "" {
		file, err := os.Create(*outPtr)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error creating output file:", err)
			os.Exit(1)
		}
		defer file.Close()
		output = file
	} else {
		output = os.Stdout
	}

	handler := lab2.NewComputeHandler(input, output)
	err := handler.Compute()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error computing expression:", err)
		os.Exit(1)
	}

}
