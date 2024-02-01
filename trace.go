package gotraceutil

import (
	"context"
)

type TracingIDGenerator func() string

var tracingIDGenerator TracingIDGenerator

func SetTracingIDGenerator(generator TracingIDGenerator) {
	tracingIDGenerator = generator
}

func Trace(ctx context.Context) context.Context {
	return context.WithValue(ctx, tracingIDKey, tracingIDGenerator())
}

func Parse(ctx context.Context) map[string]string {
	if ctx == nil {
		return nil
	}

	labels := make(map[string]string)

	for _, key := range tracingKeys {
		val := ctx.Value(key)

		if val == nil {
			labels[key] = ""
		} else {
			labels[key] = val.(string)
		}
	}

	return labels
}
