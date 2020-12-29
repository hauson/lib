package statisitcs

import (
	"testing"

	"github.com/bytom/blockcenter/test/suit"
)

func TestMedian(t *testing.T) {
	suit.TestSuit{
		{
			Desc:        "odd number",
			Args:        []float64{1.0, 3.0, 8.0, 2.0, 4.0},
			WantResults: 3.0,
		},
		{
			Desc:        "even number",
			Args:        []float64{3.0, 8.0, 2.0, 4.0},
			WantResults: 3.5,
		},
		{
			Desc:    "error",
			Args:    []float64{},
			WantErr: "inputs is empty",
		},
	}.Range(t, func(c *suit.TestCase) (interface{}, error) {
		numbers := c.Args.([]float64)
		return Median(numbers)
	})
}
