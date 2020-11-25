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

// FillOSArgs use for blockcenter, becasue it fill cfg is sick
// demo:
// go.mod
// github.com/hauson/lib/path v0.0.0
//[replace]
// github.com/hauson/lib/path => ../../../github.com/hauson/lib/path
//[src.go]
// path.FillOSArgs("config_travis.json", "blockcenter")
func FillOSArgs(relativeCfg, projRoot string) {
	if len(os.Args) == 1 {
		relativePath := RelativePath(relativeCfg, projRoot)
		os.Args = append(os.Args, relativePath)
	}
}
