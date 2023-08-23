package gotraceutil

import (
	"github.com/sirupsen/logrus"
)

type LogrusHook struct{}

func NewLogrusHook() logrus.Hook {
	return &LogrusHook{}
}

func (hook *LogrusHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook *LogrusHook) Fire(entry *logrus.Entry) error {
	for key, val := range Parse(entry.Context) {
		entry.Data[key] = val
	}

	return nil
}
