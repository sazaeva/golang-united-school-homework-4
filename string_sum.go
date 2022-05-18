package string_sum

import (
	"errors"
	"fmt"
	"strconv"
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
	n := ""
	s := []int{}
	set := 0
	output := 0
	for i, aR := range input {
		if aR >= 48 && aR <= 58 {
			n += fmt.Sprint(int(aR - 48))
			set = 1
		} else if set == 1 {
			w, err := strconv.Atoi(n)
			fmt.Println(errorNotTwoOperands)
			if err != nil {
				fmt.Println(errorEmptyInput)
			}
			fmt.Println(w)
			s = append(s, w)
			n = ""
			set = 0
		}
		if i == len(input)-1 {
			w, _ := strconv.Atoi(n)
			fmt.Println(w)
			s = append(s, w)
		}
	}
	for _, sR := range s {
		output += sR
	}

	return string(output), nil

}
