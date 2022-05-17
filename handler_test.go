package lab2

import (
	"bytes"
	. "gopkg.in/check.v1"
	"strings"
)

func (s *TestSuite) TestComputeHandler(c *C) {
	b := bytes.NewBuffer(make([]byte, 0))

	handler := ComputeHandler{
		Input:  strings.NewReader("+ 2 2"),
		Output: b,
	}
	err := handler.Compute()

	c.Assert(err, Equals, nil)
	c.Assert(b.String(), Equals, "2 + 2")
}

func (s *TestSuite) TestComputeHandlerError(c *C) {
	b := bytes.NewBuffer(make([]byte, 0))

	handler := ComputeHandler{
		Input:  strings.NewReader("23 22"),
		Output: b,
	}
	err := handler.Compute()

	c.Assert(err, NotNil)
}
