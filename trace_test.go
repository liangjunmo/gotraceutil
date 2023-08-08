package gotraceutil_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/liangjunmo/gotraceutil"
)

func TestTrace(t *testing.T) {
	traceId := "trace-id-unique-string"

	gotraceutil.SetTraceIdGenerator(func() string {
		return traceId
	})

	gotraceutil.SetTraceIdKey("TraceId")
	gotraceutil.AppendTraceKeys([]string{"ClientId"})

	ctx := gotraceutil.Trace(context.Background())
	assert.Equal(t, traceId, ctx.Value("TraceId"))
	assert.Nil(t, ctx.Value("ClientId"))

	labels := gotraceutil.Parse(ctx)
	assert.Equal(t, ctx.Value("TraceId"), labels["TraceId"])
	assert.Nil(t, labels["ClientId"])
}
