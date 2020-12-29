package statisitcs

import (
	"testing"

	"github.com/bytom/blockcenter/test/suit"
)

func TestKMax(t *testing.T) {
	type Args struct {
		Numbers []float64
		K       int
	}
	suit.TestSuit{
		{
			Desc: "k exceed error",
			Args: Args{
				Numbers: []float64{1.0},
				K:       1,
			},
			WantErr: "k 1 exceed number last index 0",
		},
		{
			Desc: "number is empty",
			Args: Args{
				Numbers: []float64{},
				K:       0,
			},
			WantErr: "numbers is empty",
		},
		{

			Desc: "k is 0",
			Args: Args{
				Numbers: []float64{1.0},
				K:       0,
			},
			WantResults: 1.0,
		},
		{
			Desc: "k is 1",
			Args: Args{
				Numbers: []float64{1.0, 2.0},
				K:       1,
			},
			WantResults: 1.0,
		},
		{
			Desc: "k is 7",
			Args: Args{
				Numbers: []float64{9.0, 4.0, 10.0, 5.0, 2.0, 3.0, 1.0, 8.0, 1.0},
				K:       7,
			},
			WantResults: 1.0,
		},
		{
			Desc: "k is 8",
			Args: Args{
				Numbers: []float64{9.0, 4.0, 10.0, 5.0, 2.0, 3.0, 1.0, 8.0, 1.0},
				K:       8,
			},
			WantResults: 1.0,
		},
	}.Range(t, func(c *suit.TestCase) (interface{}, error) {
		args := c.Args.(Args)
		return KMax(args.Numbers, args.K)
	})
}
