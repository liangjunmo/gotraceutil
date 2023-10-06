package gotraceutil_test

import (
	"context"
	"testing"

	"github.com/sirupsen/logrus"

	"github.com/liangjunmo/gotraceutil"
)

func TestLogrusHook(t *testing.T) {
	log := logrus.New()

	log.SetFormatter(&logrus.TextFormatter{
		DisableQuote:    true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	log.AddHook(
		gotraceutil.NewLogrusHook(),
	)

	ctx := context.WithValue(context.Background(), gotraceutil.DefaultTraceIDKey, "trace-id")

	log.WithContext(ctx).Error("error message with TraceID")
}
