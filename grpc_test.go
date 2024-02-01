package gotraceutil

import (
	"context"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/stretchr/testify/require"
)

func TestGRPCUnaryServerInterceptor(t *testing.T) {
	resetTracingKeys()

	tracingIDKey := "TracingID"
	tracingIDValue := "TracingValue"

	clientIDKey := "ClientID"
	clientIDValue := "ClientValue"

	SetTracingIDKey(tracingIDKey)

	SetTracingIDGenerator(func() string {
		return tracingIDValue
	})

	AppendTracingKeys([]string{clientIDKey})

	md := metadata.New(map[string]string{
		clientIDKey: clientIDValue,
	})

	ctx := metadata.NewIncomingContext(context.Background(), md)

	handler := func(ctx context.Context, req any) (any, error) {
		require.Equal(t, tracingIDValue, ctx.Value(tracingIDKey))
		require.Equal(t, clientIDValue, ctx.Value(clientIDKey))

		return nil, nil
	}

	_, err := GRPCUnaryServerInterceptor(ctx, nil, nil, handler)
	require.Nil(t, err)
}

func TestGRPCUnaryClientInterceptor(t *testing.T) {
	resetTracingKeys()

	tracingIDKey := "TracingID"
	tracingIDValue := "TracingValue"

	clientIDKey := "ClientID"
	clientIDValue := "ClientValue"

	SetTracingIDKey(tracingIDKey)

	SetTracingIDGenerator(func() string {
		return tracingIDValue
	})

	AppendTracingKeys([]string{clientIDKey})

	md := metadata.New(map[string]string{
		tracingIDKey: "",
		clientIDKey:  "",
	})

	ctx := context.WithValue(context.Background(), tracingIDKey, tracingIDValue)
	ctx = context.WithValue(ctx, clientIDKey, clientIDValue)
	ctx = metadata.NewOutgoingContext(ctx, md)

	handler := func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, opts ...grpc.CallOption) error {
		require.Equal(t, tracingIDValue, ctx.Value(tracingIDKey))
		require.Equal(t, clientIDValue, ctx.Value(clientIDKey))

		return nil
	}

	err := GRPCUnaryClientInterceptor(ctx, "", nil, nil, nil, handler, nil)
	require.Nil(t, err)
}
