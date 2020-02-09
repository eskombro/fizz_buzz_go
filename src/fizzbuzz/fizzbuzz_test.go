package fizzbuzz

import (
	"fmt"
	"testing"

	"bou.ke/monkey"
)

func TestFizzbuzz(t *testing.T) {
	monkey.Patch(AddRequest, func(params FizzBuzzParams) {})
	// Error cases
	error_tests := []FizzBuzzParams{
		FizzBuzzParams{0, 5, 15, "F", "B"},
		FizzBuzzParams{3, 0, 15, "F", "B"},
		FizzBuzzParams{3, 5, 0, "F", "B"},
		FizzBuzzParams{3, 5, 15, "", "B"},
		FizzBuzzParams{3, 5, 15, "F", ""},
	}
	error_expect := []string{"", "", "", "", ""}
	for i, test := range error_tests {
		resp := Fizzbuzz(test)
		if error_expect[i] != resp {
			t.Errorf("Error FizzBuzz.\nExpected %s\nGot      %s", error_expect[i], resp)
		} else {
			fmt.Println("Test Error case OK: ", error_tests[i])
		}
	}

	// Result tests
	tests := []FizzBuzzParams{
		FizzBuzzParams{1, 2, 3, "F", "B"},
		FizzBuzzParams{1, 3, 8, "F", "B"},
		FizzBuzzParams{3, 5, 15, "Fizz", "Buzz"},
	}
	tests_expect := []string{
		"F, FB, F",
		"F, F, FB, F, F, FB, F, F",
		"1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, FizzBuzz",
	}
	for i, test := range tests {
		resp := Fizzbuzz(test)
		if tests_expect[i] != resp {
			t.Errorf("Error FizzBuzz.\nExpected %s\nGot      %s", tests_expect[i], resp)
		} else {
			fmt.Println("Test case OK: ", tests[i])
		}
	}
}
