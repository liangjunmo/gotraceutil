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

	ctx := context.WithValue(context.Background(), tracingIDKey, tracingIDValue)
	labels := Parse(ctx)
	require.Equal(t, tracingIDValue, labels[tracingIDKey])
}
