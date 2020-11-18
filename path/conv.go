package path

import (
	"os"
	"strings"
)

// RelativePath abs path convert relative path base on projRoot
func RelativePath(path, projRoot string) string {
	curDir, _ := os.Getwd()
	var relativePath string
	var flag bool
	for _, s := range strings.Split(curDir, "/") {
		if flag {
			relativePath += "../"
		}

		if s == projRoot {
			flag = true
		}
	}

	if !flag { // not exist
		return ""
	}
	return relativePath + path
}
