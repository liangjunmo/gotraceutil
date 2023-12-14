package gotraceutil_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/liangjunmo/gotraceutil"
)

func TestGRPCUnaryServerInterceptor(t *testing.T) {
	tracingIDKey := "TracingID"
	tracingID := "tracingID"

	clientIDKey := "ClientID"
	clientID := "clientID"

	gotraceutil.SetTracingKeys([]string{tracingIDKey, clientIDKey})

	gotraceutil.SetTracingIDGenerator(func() string {
		return tracingID
	})

	ctx := context.Background()
	ctx = context.WithValue(ctx, clientIDKey, clientID)

	expected := "resp"

	handler := func(ctx context.Context, req any) (any, error) {
		assert.Equal(t, tracingID, ctx.Value(tracingIDKey))
		assert.Equal(t, clientID, ctx.Value(clientIDKey))

		return expected, nil
	}

	resp, err := gotraceutil.GRPCUnaryServerInterceptor(ctx, nil, nil, handler)
	assert.Nil(t, err)
	assert.Equal(t, expected, resp)
}
