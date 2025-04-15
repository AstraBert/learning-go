package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function
func DivideFood(fodderCalculator FodderCalculator, cows int) (float64, error) {
	totalFodder, err := fodderCalculator.FodderAmount(cows)
	if err == nil {
		feedingFactor, err := fodderCalculator.FatteningFactor()
		if err == nil {	
			foodPerCow := (feedingFactor * totalFodder) / float64(cows)
			return foodPerCow, nil
		} else {
			return 0, err
		}
	} else {
		return 0, err
	}

}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fodderCalculator FodderCalculator, cows int) (float64, error) {
	if cows > 0 {
		foodPerCow, err := DivideFood(fodderCalculator, cows)
		return foodPerCow, err
	} else {
		return 0, errors.New("invalid number of cows")
	}
}

type InvalidCowsError struct {
	cowsNumber int
	message string
}

func newInvalidCowsError(cowsNumber int) InvalidCowsError {
	return InvalidCowsError{
		cowsNumber: cowsNumber,
		message: "",
	}
}

func (e *InvalidCowsError) Error() string {
	if e.cowsNumber < 0 {
		e.message = "there are no negative cows"
	} else {
		e.message = "no cows don't need food"
	}
	return fmt.Sprintf("%d cows are invalid: %s", e.cowsNumber, e.message)
}  

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(cows int) (error) {
	if cows > 0 {
		return nil
	} else {
		err := newInvalidCowsError(cows)
		return &err
	}
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
