package gotraceutil

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	resetTracingKeys()

	tracingIDKey := "TracingID"
	tracingIDValue := "TracingValue"

	ctx := context.Background()
	ctx = context.WithValue(ctx, tracingIDKey, tracingIDValue)

	labels := Parse(ctx)
	require.Equal(t, tracingIDValue, labels[tracingIDKey])
}
