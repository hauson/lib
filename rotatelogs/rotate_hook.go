package rotatelogs

import (
	"sync"
	"time"
	"path/filepath"

	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

const (
	day          time.Duration = 24 * time.Hour
	rotationTime time.Duration = 1 * day
	maxAge       time.Duration = 7 * day
)

var defaultFormatter = &logrus.TextFormatter{DisableColors: true}

type rotateHook struct {
	logPath string
	lock    *sync.Mutex
}

func newRotateHook(logPath string) *rotateHook {
	return &rotateHook{
		lock:    new(sync.Mutex),
		logPath: logPath,
	}
}

// Write a log line to an io.Writer.
func (hook *rotateHook) ioWrite(entry *logrus.Entry) error {
	module := "general"
	if data, ok := entry.Data["module"]; ok {
		module = data.(string)
	}

	logPath := filepath.Join(hook.logPath, module)
	writer, err := rotatelogs.New(
		logPath+".%Y%m%d%H%M%S",
		rotatelogs.WithLinkName(logPath+".log"),
		rotatelogs.WithMaxAge(maxAge),
		rotatelogs.WithRotationTime(rotationTime),
	)
	if err != nil {
		return err
	}

	msg, err := defaultFormatter.Format(entry)
	if err != nil {
		return err
	}

	if _, err = writer.Write(msg); err != nil {
		return err
	}

	return writer.Close()
}

//Fire write to file
func (hook *rotateHook) Fire(entry *logrus.Entry) error {
	hook.lock.Lock()
	defer hook.lock.Unlock()

	return hook.ioWrite(entry)
}

// Levels returns configured log levels.
func (hook *rotateHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
