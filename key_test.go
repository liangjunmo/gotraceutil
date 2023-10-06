package gotraceutil_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/liangjunmo/gotraceutil"
)

func TestKey(t *testing.T) {
	keys := gotraceutil.GetTraceKeys()
	assert.Equal(t, gotraceutil.DefaultTraceIDKey, keys[0])

	traceIDKey := "RequestID"
	gotraceutil.SetTraceIDKey(traceIDKey)
	keys = gotraceutil.GetTraceKeys()
	assert.Equal(t, traceIDKey, keys[0])

	clientIDKey := "clientID"
	gotraceutil.ResetTraceKeys()
	gotraceutil.AppendTraceKeys([]string{clientIDKey})
	keys = gotraceutil.GetTraceKeys()
	assert.Equal(t, gotraceutil.DefaultTraceIDKey, keys[0])
	assert.Equal(t, clientIDKey, keys[1])
}
