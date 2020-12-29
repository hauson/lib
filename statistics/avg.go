package statisitcs

import (
	"errors"
)

// Avg calc avg of float64 numbers
func Avg(numbers []float64) (float64, error) {
	if len(numbers) == 0 {
		return 0.0, errors.New("numbers is empty")
	}

	var sum float64
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers)), nil
}
