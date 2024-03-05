package gotraceutil

import (
	"bytes"
	"context"
	"encoding/json"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestLogrusHook(t *testing.T) {
	resetTracingKeys()

	tracingIDKey := "TracingID"
	tracingIDValue := "TracingValue"

	SetTracingIDKey(tracingIDKey)

	SetTracingIDGenerator(func() string {
		return tracingIDValue
	})

	ctx := Trace(context.Background())

	var (
		buffer bytes.Buffer
		fields logrus.Fields
	)

	log := logrus.New()

	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(&buffer)

	log.AddHook(NewLogrusHook())

	log.WithContext(ctx).Error("message")

	err := json.Unmarshal(buffer.Bytes(), &fields)
	require.Nil(t, err)
	require.Equal(t, tracingIDValue, fields[tracingIDKey])
}
