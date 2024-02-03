package gotraceutil

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTrace(t *testing.T) {
	resetTracingKeys()

	tracingIDKey := "TracingID"
	tracingIDValue := "TracingValue"

	SetTracingIDKey(tracingIDKey)

	SetTracingIDGenerator(func() string {
		return tracingIDValue
	})

	ctx := Trace(context.Background())
	require.Equal(t, tracingIDValue, ctx.Value(tracingIDKey))
}
