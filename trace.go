package gotraceutil

import (
	"context"

	uuid "github.com/satori/go.uuid"
)

var DefaultTraceIDGenerator = func() string {
	return uuid.NewV4().String()
}

var traceIDGenerator func() string

func SetTraceIDGenerator(fn func() string) {
	traceIDGenerator = fn
}

func Trace(ctx context.Context) context.Context {
	return context.WithValue(ctx, traceIDKey, traceIDGenerator())
}

func Parse(ctx context.Context) map[string]interface{} {
	if ctx == nil {
		return nil
	}

	var labels map[string]interface{}

	for _, key := range traceKeys {
		val := ctx.Value(key)
		if val == nil {
			continue
		}

		if labels == nil {
			labels = make(map[string]interface{})
		}

		labels[key] = val
	}

	return labels
}
