package print

import (
	"fmt"
	"regexp"
	"runtime"
)

var goSrcRegexp = regexp.MustCompile(`hauson/lib(@.*)?/.*.go`)
var goTestRegexp = regexp.MustCompile(`hauson/lib(@.*)?/.*test.go`)

// FileLineNum file and line name
func FileLineNum() string {
	for i := 2; i < 15; i++ {
		_, file, line, ok := runtime.Caller(i)
		if ok && (!goSrcRegexp.MatchString(file) || goTestRegexp.MatchString(file)) {
			return fmt.Sprintf("%v:%v", file, line)
		}
	}
	return ""
}
