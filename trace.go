package gotraceutil

import (
	"context"

	uuid "github.com/satori/go.uuid"
)

var tracingIDGenerator = func() string {
	return uuid.NewV4().String()
}

func SetTracingIDGenerator(fn func() string) {
	tracingIDGenerator = fn
}

func Trace(ctx context.Context) context.Context {
	return context.WithValue(ctx, tracingKeys[0], tracingIDGenerator())
}

func Parse(ctx context.Context) map[string]interface{} {
	if ctx == nil {
		return nil
	}

	var labels map[string]interface{}

	for _, key := range tracingKeys {
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
