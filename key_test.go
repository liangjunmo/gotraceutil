package gotraceutil_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/liangjunmo/gotraceutil"
)

func TestKeys(t *testing.T) {
	keys := gotraceutil.GetTraceKeys()
	assert.Equal(t, "TraceId", keys[0])

	gotraceutil.SetTraceIDKey("RequestId")
	keys = gotraceutil.GetTraceKeys()
	assert.Equal(t, "RequestId", keys[0])

	gotraceutil.SetTraceIDKey("TraceId")
	gotraceutil.AppendTraceKeys([]string{"ClientId"})
	keys = gotraceutil.GetTraceKeys()
	assert.Equal(t, "TraceId", keys[0])
	assert.Equal(t, "ClientId", keys[1])
}
