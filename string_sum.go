package string_sum

import (
	"errors"
	"fmt"
	"regexp"
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

func StringSum(input string) (string, error) {

	input = strings.ReplaceAll(input, " ", "")
	if input == "" {
		return "", fmt.Errorf("%w", errorEmptyInput)
	}

	reg := regexp.MustCompile(`[+-]?\d+[^+-]`)
	slice := reg.FindAllString(input, -1)

	if len(slice) != 2 {
		return "", fmt.Errorf("%w", errorNotTwoOperands)
	}

	operand1, err := strconv.Atoi(strings.TrimSpace(slice[0]))
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	operand2, err := strconv.Atoi(strings.TrimSpace(slice[1]))
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	return strconv.Itoa(operand1 + operand2), nil
}
