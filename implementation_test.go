package lab2

import (
	"fmt"
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type TestSuite struct{}

var _ = Suite(&TestSuite{})

func (s *TestSuite) TestPrefixToInfix(c *C) {
	examples := map[string]string{
		"+ 5 * - 4 2 3":                    "5 + (4 - 2) * 3",
		"* 7 - 1 * 13 4":                   "7 * (1 - 13 * 4)",
		"- 4 3 2":                          "invalid input expression",
		"+ 11 + 22 + 33 + 4 + 5 + 6 + 7 8": "11 + 22 + 33 + 4 + 5 + 6 + 7 + 8",
		"/ + * 14 - 15 7 * 3 17 + 10 1":    "(14 * (15 - 7) + 3 * 17) / (10 + 1)",
		"Some text":                        "invalid input expression",
		"- - + 777 - 1901.2021 - 12":       "invalid input expression",
		"":                                 "invalid input expression",
	}

	for prefix, infix := range examples {
		res, err := PrefixToInfix(prefix)
		if err != nil {
			c.Assert(err, ErrorMatches, infix)
		} else {
			c.Assert(res, Equals, infix)
		}
	}
}

func ExamplePrefixToInfix() {
	res, err := PrefixToInfix("14 11 + 1 ^")
	if err != nil {
		panic(err)
	} else {
		fmt.Println(res)
	}
}
