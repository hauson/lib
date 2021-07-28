package statisitcs

import (
	"testing"

	"github.com/hauson/lib/testsuit"
)

func TestMedianAvg(t *testing.T) {
	testsuit.TestSuit{
		{
			Desc:        "normal",
			Args:        []float64{1.0, 3.0, 8.0, 2.0, 4.0},
			WantResults: 3.0,
		},
		{
			Desc:    "null",
			Args:    []float64{},
			WantErr: "numbers is empty",
		},
	}.Range(t, func(c *testsuit.TestCase) (interface{}, error) {
		numbers := c.Args.([]float64)
		return MedianAvg(numbers)
	})
}
