package statisitcs

import "math"

const threshold = 3.5

//MADFilter median absolute deviation filter
func MADFilter(numbers []float64) ([]float64, error) {
	median, err := Median(numbers)
	if err != nil {
		return nil, err
	}

	diffs := make([]float64, len(numbers))
	for i, f := range numbers {
		diffs[i] = math.Abs(f - median)
	}

	filterNumbers := make([]float64, 0, len(numbers))
	medianDiff, err := Median(diffs)
	if err != nil {
		return nil, err
	}

	for i, diff := range diffs {
		if diff > threshold*medianDiff {
			continue
		}
		filterNumbers = append(filterNumbers, numbers[i])
	}
	return filterNumbers, nil
}
