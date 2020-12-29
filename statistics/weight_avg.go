package statisitcs

import "errors"

// WeightNumber number with weight
type WeightNumber struct {
	Value  float64
	Weight float64
}

// WeightAvg weight Avg
func WeightAvg(numbers []*WeightNumber) (float64, error) {
	if len(numbers) == 0 {
		return 0.0, errors.New("numbers is empty")
	}

	var totalValue, totalWight float64
	for _, number := range numbers {
		if number.Weight < 0 {
			return 0, errors.New("weight less zero")
		}

		totalWight += number.Weight
		totalValue += number.Weight * number.Value
	}

	return totalValue / totalWight, nil
}
