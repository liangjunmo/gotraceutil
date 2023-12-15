package gotraceutil_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/liangjunmo/gotraceutil"
)

func TestTrace(t *testing.T) {
	tracingIDKey := "TracingID"
	tracingIDVal := "tracing-id"

	gotraceutil.SetTracingKeys([]string{tracingIDKey})

	gotraceutil.SetTracingIDGenerator(func() string {
		return tracingIDVal
	})

	ctx := gotraceutil.Trace(context.Background())
	assert.Equal(t, tracingIDVal, ctx.Value(tracingIDKey))

	labels := gotraceutil.Parse(ctx)
	assert.Equal(t, tracingIDVal, labels[tracingIDKey])
}
