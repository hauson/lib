package statisitcs

import (
	"testing"

	"github.com/bytom/blockcenter/test/suit"
)

func TestMADFilter(t *testing.T) {
	suit.TestSuit{
		{
			Desc:        "normal",
			Args:        []float64{1.0, 1.0001, 1.0002, 1.0003},
			WantResults: []float64{1.0, 1.0001, 1.0002, 1.0003},
			WantErr:     "",
		},
		{
			Desc:        "filter",
			Args:        []float64{1.0, 1.0001, 1.0002, 1.000500000000001},
			WantResults: []float64{1.0, 1.0001, 1.0002},
			WantErr:     "",
		},
		{
			Desc:        "null",
			Args:        []float64{},
			WantResults: []float64{},
			WantErr:     "inputs is empty",
		},
		{
			Desc:        "same numbers",
			Args:        []float64{1, 1, 1},
			WantResults: []float64{1, 1, 1},
		},
	}.Range(t, func(c *suit.TestCase) (interface{}, error) {
		numbers := c.Args.([]float64)
		return MADFilter(numbers)
	})
}
