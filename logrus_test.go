package gotraceutil

import (
	"context"

	"github.com/sirupsen/logrus"
)

func ExampleLogrusHook() {
	resetTracingKeys()

	tracingIDKey := "TracingID"
	tracingIDValue := "TracingValue"

	SetTracingIDKey(tracingIDKey)

	SetTracingIDGenerator(func() string {
		return tracingIDValue
	})

	ctx := Trace(context.Background())

	log := logrus.New()

	log.SetFormatter(&logrus.TextFormatter{
		DisableQuote:    true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	log.AddHook(NewLogrusHook())

	log.WithContext(ctx).Error("message")
}
