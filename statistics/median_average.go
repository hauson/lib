package statisitcs

import "errors"

// MedianAvg move a max, and min, then calc avg
func MedianAvg(numbers []float64) (float64, error) {
	switch len(numbers) {
	case 0:
		return 0.0, errors.New("numbers is empty")
	case 1, 2:
		return Avg(numbers)
	default:
		min, max, sum := numbers[0], numbers[0], 0.0
		for _, number := range numbers {
			if max < number {
				max = number
			}
			if min > number {
				min = number
			}
			sum += number
		}

		return (sum - min - max) / float64((len(numbers) - 2)), nil
	}
}
