package statisitcs

import (
	"testing"

	"github.com/hauson/lib/testsuit"
)

func TestAvg(t *testing.T) {
	testsuit.TestSuit{
		{
			Desc:        "avg",
			Args:        []float64{1.0, 3.0, 8.0, 2.0, 4.0},
			WantResults: 3.6,
			WantErr:     "",
		},
		{
			Desc:    "null",
			Args:    []float64{},
			WantErr: "numbers is empty",
		},
	}.Range(t, func(c *testsuit.TestCase) (interface{}, error) {
		numbers := c.Args.([]float64)
		return Avg(numbers)
	})
}
