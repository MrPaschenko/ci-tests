package main

import (
	"flag"
	lab2 "github.com/MrPaschenko/ci-tests"
	"io"
	"os"
	"strings"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFromFile   = flag.String("f", "", "Expression from file")
	outputToFile    = flag.String("o", "", "Save results to file")
)

func main() {
	flag.Parse()

	var input io.Reader = nil
	var output = os.Stdout

	if *inputExpression != "" {
		input = strings.NewReader(*inputExpression)
	}

	if *inputFromFile != "" {
		f, err := os.Open(*inputFromFile)
		if err != nil {
			os.Stderr.WriteString("Error")
		}
		defer f.Close()
		input = f
	}

	if *outputToFile != "" {
		f, err := os.Create(*outputToFile)
		if err != nil {
			os.Stderr.WriteString("Error")
		}
		defer f.Close()
		output = f
	}

	if input == nil {
		os.Stderr.WriteString("No stdIn defined")
		return
	}

	handler := &lab2.ComputeHandler{
		Input:  input,
		Output: output,
	}
	err := handler.Compute()
	if err != nil {
		println(err) // print err to stderr
	}
}
