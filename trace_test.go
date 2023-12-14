package gotraceutil_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/liangjunmo/gotraceutil"
)

func TestTrace(t *testing.T) {
	tracingIDKey := "TracingID"
	tracingID := "tracingID"

	gotraceutil.SetTracingKeys([]string{tracingIDKey})

	gotraceutil.SetTracingIDGenerator(func() string {
		return tracingID
	})

	ctx := gotraceutil.Trace(context.Background())
	assert.Equal(t, tracingID, ctx.Value(tracingIDKey))

	labels := gotraceutil.Parse(ctx)
	assert.Equal(t, tracingID, labels[tracingIDKey])
}
