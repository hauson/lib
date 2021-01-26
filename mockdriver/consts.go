package mockdriver

// CURDMode curd mode
type CURDMode string

// String to string
func (c CURDMode) String() string {
	return string(c)
}

const (
	CURDUnkown CURDMode = "CURD_UNKOWN"
	CURDInsert CURDMode = "INSERT"
	CURDSelect CURDMode = "SELECT"
	CURDUpdate CURDMode = "UPDATE"
	CURDDelete CURDMode = "DELETE"
)


