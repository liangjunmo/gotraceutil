package gotraceutil_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/liangjunmo/gotraceutil"
)

func TestTrace(t *testing.T) {
	traceID := "trace-id"

	gotraceutil.SetTraceIDGenerator(func() string {
		return traceID
	})

	ctx := gotraceutil.Trace(context.Background())
	assert.Equal(t, traceID, ctx.Value(gotraceutil.DefaultTraceIDKey))

	labels := gotraceutil.Parse(ctx)
	assert.Equal(t, traceID, labels[gotraceutil.DefaultTraceIDKey])
}
