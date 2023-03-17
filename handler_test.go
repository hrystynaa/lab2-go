package lab2

import (
	"bytes"
	"strings"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type ComputeHandlerSuite struct{}

var _ = Suite(&ComputeHandlerSuite{})

func (s *ComputeHandlerSuite) TestComputeHandlerValid(c *C) {
	input := strings.NewReader("+ 5 * - 4 2 3")
	output := bytes.NewBuffer(nil)

	ch := NewComputeHandler(input, output)

	err := ch.Compute()

	c.Check(err, IsNil)
	c.Check(output.String(), Equals, "4 2 - 3 * 5 +\n")

}

func (s *ComputeHandlerSuite) TestComputeHandlerInvalid(c *C) {
	input := strings.NewReader("? + 4 * 8 9 6")
	output := bytes.NewBuffer(nil)

	ch := NewComputeHandler(input, output)

	err := ch.Compute()

	c.Check(err, NotNil)
}
