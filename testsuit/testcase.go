package testsuit

// TestCase test case
type TestCase struct {
	Desc        string
	Args        interface{}
	WantResults interface{}
	WantErr     string
}

func (c *TestCase) IsWantErr(err error) bool {
	var errInfo string
	if err != nil {
		errInfo = err.Error()
	}
	return errInfo == c.WantErr
}

func (c *TestCase) DiffWantResults(got interface{}) string {
	return Diff(c.WantResults, got)
}
