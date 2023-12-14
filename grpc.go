package gotraceutil

import (
	"context"

	"google.golang.org/grpc"
)

func GRPCUnaryServerInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	for _, key := range tracingKeys {
		val := ctx.Value(key)
		if val != nil {
			continue
		}

		ctx = context.WithValue(ctx, key, "")
	}

	tracingID := ctx.Value(tracingKeys[0])
	if tracingID == "" {
		ctx = Trace(ctx)
	}

	return handler(ctx, req)
}
