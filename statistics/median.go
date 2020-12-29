package statisitcs

import (
	"errors"
)

// Median return float64 numbers median
func Median(numbers []float64) (float64, error) {
	if len(numbers) == 0 {
		return 0, errors.New("inputs is empty")
	}

	middle := len(numbers) / 2
	// length is odd
	if len(numbers)%2 == 0 {
		leftMedian, err := KMax(numbers, middle-1)
		if err != nil {
			return 0.0, err
		}

		rightMedian, err := KMax(numbers, middle)
		if err != nil {
			return 0.0, err
		}
		return (leftMedian + rightMedian) / 2, nil
	}

	// length is even
	return KMax(numbers, middle)
}
