package statisitcs

import (
	"testing"

	"github.com/hauson/lib/testsuit"
)

func TestWeightAvg(t *testing.T) {
	testsuit.TestSuit{
		{
			Desc: "error",
			Args: []*WeightNumber{
				{
					Value:  3,
					Weight: 1,
				},
				{
					Value:  5,
					Weight: -2,
				},
			},
			WantErr: "weight less zero",
		},
		{
			Desc: "error",
			Args: []*WeightNumber{
				{
					Value:  3,
					Weight: 1,
				},
				{
					Value:  5,
					Weight: -2,
				},
			},
			WantErr: "weight less zero",
		},
		{
			Desc:    "null",
			Args:    []*WeightNumber{},
			WantErr: "numbers is empty",
		},
		{
			Desc: "diff weight",
			Args: []*WeightNumber{
				{
					Value:  2,
					Weight: 2,
				},
				{
					Value:  5,
					Weight: 1,
				},
			},
			WantResults: 3.0,
		},
	}.Range(t, func(c *testsuit.TestCase) (interface{}, error) {
		numbers := c.Args.([]*WeightNumber)
		return WeightAvg(numbers)
	})
}
