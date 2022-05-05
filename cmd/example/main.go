package main

import (
	"flag"
	"fmt"
	lab2 "github.com/roman-mazur/architecture-lab-2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	// TODO: Add other flags support for input and output configuration.
)

func isOperator(x) {
	switch x {
		case "+":
		case "-":
		case "/":
		case "*":
		case "^":
		case "%":
			return true
	}
	return false
}

// Convert prefix to infix expression
func prefixToInfix(prefixExpr) {

	// length of expression	
	length := prefixExpr.size()	

	stack = make([]int, length)

	// reading from right to left
	for i:= length-1; i >= 0; i-- {
		// check if symbol is operator
		if isOperator(prefixExpr[i]) {
			//var op1 = s.
		}
	}
}

func main() {
	flag.Parse()

	// TODO: Change this to accept input from the command line arguments as described in the task and
	//       output the results using the ComputeHandler instance.
	//       handler := &lab2.ComputeHandler{
	//           Input: {construct io.Reader according the command line parameters},
	//           Output: {construct io.Writer according the command line parameters},
	//       }
	//       err := handler.Compute()

	res, _ := lab2.PrefixToPostfix("+ 2 2")
	fmt.Println(res)
}
