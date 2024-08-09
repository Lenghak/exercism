package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function
func DivideFood(f FodderCalculator, cows int) (float64, error) {

	fodder, fodderErr := f.FodderAmount(cows)
	factor, factorErr := f.FatteningFactor()

	if fodderErr != nil {
		return 0, fodderErr
	}

	if factorErr != nil {
		return 0, factorErr
	}

	return factor * fodder / float64(cows), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(f FodderCalculator, cows int) (float64, error) {
	if cows > 0 {
		return DivideFood(f, cows)
	}

	return 0, errors.New("invalid number of cows")
}

type InvalidCowsError struct {
	cows    int
	message string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.message)
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return &InvalidCowsError{cows: cows, message: "there are no negative cows"}
	}

	if cows == 0 {
		return &InvalidCowsError{cows: cows, message: "no cows don't need food"}
	}
	return nil
}
