package lab2

import (
	"bufio"
	"fmt"
	"io"
)

type ComputeHandler struct {
	input  io.Reader
	output io.Writer
}

func NewComputeHandler(input io.Reader, output io.Writer) *ComputeHandler {
	return &ComputeHandler{input, output}
}

func (ch *ComputeHandler) Compute() error {
	scanner := bufio.NewScanner(ch.input)
	scanner.Scan()
	expr := scanner.Text()

	result, err := PrefixToPostfix(expr)
	if err != nil {
		return fmt.Errorf("error computing expression: %v", err)
	}

	_, err = fmt.Fprintln(ch.output, result)
	if err != nil {
		return fmt.Errorf("error writing result: %v", err)
	}

	return nil
}
