package rotatelogs

import (
	"fmt"
	"os"
	"strings"
	"io/ioutil"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

//InitLogFile init logrus with hook
func InitLogFile(logPath string) error {
	if err := clearLockFiles(logPath); err != nil {
		return err
	}

	logrus.AddHook(newRotateHook(logPath))
	logrus.SetOutput(ioutil.Discard)
	fmt.Printf("all logs are output in the %s directory\n", logPath)
	return nil
}

func clearLockFiles(logPath string) error {
	files, err := ioutil.ReadDir(logPath)
	if os.IsNotExist(err) {
		return nil
	} else if err != nil {
		return err
	}

	for _, file := range files {
		if ok := strings.HasSuffix(file.Name(), "_lock"); ok {
			if err := os.Remove(filepath.Join(logPath, file.Name())); err != nil {
				return err
			}
		}
	}
	return nil
}
