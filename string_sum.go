package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	input = strings.ReplaceAll(input, " ", "")

	if input == "" || input == "-" || input == "+" {
		if err != nil {
			return input, fmt.Errorf("%w", errorEmptyInput)
		}
	}

	var container string
	var a, b, op int
	n := 1

	for i, v := range input {
		if i == 0 && v == 43 || v > 57 || len(input) > 1 {
			if err != nil {
				return "", fmt.Errorf("%w", errorNotTwoOperands)
			}
		}

		if i == 0 && v == 45 { // first character is minus
			n = -1 // first operand is negative
			continue
		}

		if v == 45 || v == 43 { // minus or plus
			if op > 0 {
				if err != nil {
					return "", fmt.Errorf("%w", errorNotTwoOperands)
				}
			}
			op = int(v)
			a, _ = strconv.Atoi(container)
			container = ""
			continue
		}

		container += string(v)
	}

	if op == 0 {
		if err != nil {
			return "", fmt.Errorf("%w", errorNotTwoOperands)
		}
	}

	b, _ = strconv.Atoi(container)

	a *= n

	switch op {
	case 45:
		output = strconv.Itoa(a - b)
	case 43:
		output = strconv.Itoa(a + b)
	}

	return output, nil
}
