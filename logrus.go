package gotraceutil

import (
	"github.com/sirupsen/logrus"
)

type logrusHook struct{}

func NewLogrusHook() logrus.Hook {
	return &logrusHook{}
}

func (hook *logrusHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook *logrusHook) Fire(entry *logrus.Entry) error {
	for key, val := range Parse(entry.Context) {
		entry.Data[key] = val
	}

	return nil
}
