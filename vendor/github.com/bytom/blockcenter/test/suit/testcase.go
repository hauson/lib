package suit

import (
	"fmt"
	"testing"
)

// TestSuit represent TestCases
type TestSuit []*TestCase

// Range range test case, and do every test case
func (s TestSuit) Range(t *testing.T, do func(*TestCase) (interface{}, error)) {
	for i, c := range s {
		if c.NotTest {
			continue
		}

		desc := fmt.Sprintf("[testcase:%d] desc:%s", i, c.Desc, )
		result, err := do(c)
		if diff := c.DiffErr(err); diff != "" {
			t.Fatalf("%s, want and got err diff: %s", desc, diff)
		}

		if err != nil {
			continue
		}

		if diff := c.DiffResult(result, c.IgnoreFileds...); diff != "" {
			t.Fatalf("%s, want and got result diff: %s", desc, diff)
		}
	}
}

// TestCase test case
type TestCase struct {
	NotTest      bool
	Desc         string
	Args         interface{}
	WantResults  interface{}
	WantErr      string
	IgnoreFileds []string
}

// DiffErr return diff want and got err,  can compatibility nil
func (c *TestCase) DiffErr(err error) string {
	var got string
	if err != nil {
		got = err.Error()
	}

	if got == c.WantErr {
		return ""
	}

	return fmt.Sprintf("want:{%s}, got:{%s}", c.WantErr, got)
}

// DiffResult cmp  want and got diff, return empty string when equal
func (c *TestCase) DiffResult(got interface{}, ignoreFields ...string) string {
	return Diff(c.WantResults, got, ignoreFields...)
}
