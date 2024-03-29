package gotraceutil

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetTracingIDKey(t *testing.T) {
	resetTracingKeys()
	require.Equal(t, DefaultTracingIDKey, GetTracingIDKey())
}

func TestSetTracingIDKey(t *testing.T) {
	resetTracingKeys()

	tracingIDKey := "ContextID"
	SetTracingIDKey(tracingIDKey)

	require.Equal(t, tracingIDKey, GetTracingIDKey())
}

func TestGetTracingKeys(t *testing.T) {
	resetTracingKeys()
	require.Equal(t, []string{DefaultTracingIDKey}, GetTracingKeys())
}

func TestAppendTracingKeys(t *testing.T) {
	resetTracingKeys()

	clientIDKey := "ClientID"
	AppendTracingKeys([]string{clientIDKey})

	require.Equal(t, []string{DefaultTracingIDKey, clientIDKey}, GetTracingKeys())
}
