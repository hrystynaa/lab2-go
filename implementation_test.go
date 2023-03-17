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
		{"complex", "+ / ^ * + 6 7 8 - 8 9 6 * / 2 3 4", "6 7 + 8 * 8 9 - ^ 6 / 2 3 / 4 * +", nil},
		{"empty", "", "", fmt.Errorf("invalid expression: ")},
		{"not enouph numbers", "+ 4", "", fmt.Errorf("invalid expression: + 4")},
		{"postfix", "6 8 ^", "", fmt.Errorf("invalid expression: 6 8 ^")},
		{"too many operators", "- + 5 * - 4 2 3", "", fmt.Errorf("invalid expression: - + 5 * - 4 2 3")},
		{"invalid token", "? + 4 * 8 9 6", "", fmt.Errorf("invalid token: ?")},
	}

	for _, test := range tests {

		c.Logf("CASE: %s", test.name)

		output, err := PrefixToPostfix(test.input)

		status := c.Check(output, Equals, test.output)
		status = status && c.Check(err, DeepEquals, test.err)

		if status {
			c.Logf("PASSED\n")
		} else {
			c.Logf("FAILED\n")
		}
	}
}

// Example demonstrates how to use PrefixToPostfix function.
func ExamplePrefixToPostfix() {
	res, _ := PrefixToPostfix("- + 2 2 * 9 2")
	fmt.Println(res)

	// Output:
	// 2 2 + 9 2 * -
}
