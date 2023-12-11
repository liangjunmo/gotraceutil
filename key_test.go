package gotraceutil_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/liangjunmo/gotraceutil"
)

func TestKey(t *testing.T) {
	tracingKeys := []string{"TracingID", "ClientID"}
	gotraceutil.SetTracingKeys(tracingKeys)
	keys := gotraceutil.GetTracingKeys()
	assert.Equal(t, tracingKeys[0], keys[0])
	assert.Equal(t, tracingKeys[1], keys[1])
}
