package gotraceutil

import (
	"context"
)

var (
	generateTraceId func() string
)

func SetTraceIdGenerator(fn func() string) {
	generateTraceId = fn
}

func Trace(ctx context.Context) context.Context {
	return context.WithValue(ctx, traceIdKey, generateTraceId())
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
