package gotraceutil_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/liangjunmo/gotraceutil"
)

func TestGRPCUnaryServerInterceptor(t *testing.T) {
	tracingIDKey := "TracingID"
	tracingIDVal := "tracing-id"

	clientIDKey := "ClientID"
	clientIDVal := "client-id"

	gotraceutil.SetTracingKeys([]string{tracingIDKey, clientIDKey})

	gotraceutil.SetTracingIDGenerator(func() string {
		return tracingIDVal
	})

	md := metadata.New(map[string]string{
		clientIDKey: clientIDVal,
	})
	ctx := metadata.NewIncomingContext(context.Background(), md)

	handler := func(ctx context.Context, req any) (any, error) {
		assert.Equal(t, tracingIDVal, ctx.Value(tracingIDKey))
		assert.Equal(t, clientIDVal, ctx.Value(clientIDKey))

		return nil, nil
	}

	_, err := gotraceutil.GRPCUnaryServerInterceptor(ctx, nil, nil, handler)
	assert.Nil(t, err)
}

func TestGRPCUnaryClientInterceptor(t *testing.T) {
	tracingIDKey := "TracingID"
	tracingIDVal := "tracing-id"

	clientIDKey := "ClientID"
	clientIDVal := "client-id"

	gotraceutil.SetTracingKeys([]string{tracingIDKey, clientIDKey})

	ctx := context.Background()
	ctx = context.WithValue(ctx, tracingIDKey, tracingIDVal)
	ctx = context.WithValue(ctx, clientIDKey, clientIDVal)

	md := metadata.New(map[string]string{
		tracingIDKey: "",
		clientIDKey:  "",
	})
	ctx = metadata.NewOutgoingContext(ctx, md)

	handler := func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, opts ...grpc.CallOption) error {
		assert.Equal(t, tracingIDVal, ctx.Value(tracingIDKey))
		assert.Equal(t, clientIDVal, ctx.Value(clientIDKey))

		return nil
	}

	err := gotraceutil.GRPCUnaryClientInterceptor(ctx, "", nil, nil, nil, handler, nil)
	assert.Nil(t, err)
}
