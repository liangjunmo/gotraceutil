package gotraceutil

import (
	"context"

	uuid "github.com/satori/go.uuid"
)

var traceIdGenerator = func() string {
	return uuid.NewV4().String()
}

func SetTraceIdGenerator(fn func() string) {
	traceIdGenerator = fn
}

func Trace(ctx context.Context) context.Context {
	return context.WithValue(ctx, traceIdKey, traceIdGenerator())
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
