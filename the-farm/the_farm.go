package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	cows int
}

func (sne *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", sne.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	if cows == 0 {
		return 0.0, errors.New("Division by zero")
	}
	if cows < 0 {
		return 0.0, &SillyNephewError{cows: cows}
	}
	weight, err := weightFodder.FodderAmount()
	if err == ErrScaleMalfunction {
		switch {
		case weight < 0:
			return 0.0, errors.New("Negative fodder")
		default:
			return (weight * 2) / float64(cows), nil
		}
	} else if err != nil {
		return 0.0, err
	} else {
		switch {
		case weight < 0:
			return 0.0, errors.New("Negative fodder")
		default:
			return weight / float64(cows), nil
		}
	}
}
