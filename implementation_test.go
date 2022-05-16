package lab2

import (
	"fmt"
	_ "gopkg.in/check.v1"
	"testing"
)

func ImplementationTest(t *testing.T) { TestingT(t) }

type TestSuite struct{}

func (s *TestSuite) prefixToinfixTest(c *C) {
	examples := map[string]string{
		"+ 5 * - 4 2 3":                    "5 + (4 - 2) * 3",
		"* 7 - 1 * 13 4":                   "7 * (1 - 13 * 4)",
		"- 4 3 2":                          "too many operands",
		"+ 11 + 22 + 33 + 4 + 5 + 6 + 7 8": "11 + 22 + 33 + 4 + 5 + 6 + 7 + 8",
		"/ + * 14 - 15 7 * 3 17 + 10 1":    "(14 * (15 - 7) + 3 * 17) / (10 + 1)",
		"Some text":                        "invalid input expression",
		"- - + 777 - 1901.2021 - 12":       "too many operators",
		"":                                 "invalid input expression",
	}

	for prefix, infix := range examples {
		res, err := prefixToinfix(prefix)
		if err != nil {
			c.Assert(err, ErrorMatches, infix)
		} else {
			c.Assert(res, Equals, infix)
		}
	}
}

func ExamplePostToIn() {
	res, err := prefixToinfix("14 11 + 1 ^")
	if err != nil {
		panic(err)
	} else {
		fmt.Println(res)
	}

	// expected output:
	// (14 + 11) ^ 1
}
