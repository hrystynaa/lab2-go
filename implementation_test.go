package lab2

import (
	"fmt"
	"testing"

	. "gopkg.in/check.v1"
)

func TestPrefixToPostfix(t *testing.T) { TestingT(t) }

type PrefixToPostfixSuite struct{}

var _ = Suite(&PrefixToPostfixSuite{})

func (s *PrefixToPostfixSuite) TestPrefixToPostfix(c *C) {
	tests := []struct {
		name   string
		input  string
		output string
		err    error
	}{
		{"simple", "+ 1 2", "1 2 +", nil},
		{"middle", "+ 5 * - 4 2 3", "4 2 - 3 * 5 +", nil},
	}

	for _, test := range tests {
		output, err := PrefixToPostfix(test.input)

		c.Check(output, Equals, test.output)
		c.Check(err, DeepEquals, test.err)
	}
}

func ExamplePrefixToPostfix() {
	res, _ := PrefixToPostfix("- + 2 2 * 9 2")
	fmt.Println(res)

	// Output:
	// 2 2 + 9 2 * -
}
