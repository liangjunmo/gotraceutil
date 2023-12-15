package gotraceutil

import (
	"context"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func GRPCUnaryServerInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	md, ok := metadata.FromIncomingContext(ctx)

	for _, key := range tracingKeys {
		if !ok {
			ctx = context.WithValue(ctx, key, "")
		} else {
			vals, ok := md[strings.ToLower(key)]
			if !ok {
				ctx = context.WithValue(ctx, key, "")
			} else if len(vals) == 0 {
				ctx = context.WithValue(ctx, key, "")
			} else {
				ctx = context.WithValue(ctx, key, vals[0])
			}
		}
	}

	if ctx.Value(tracingIDKey) == "" {
		ctx = Trace(ctx)
	}

	return handler(ctx, req)
}

func GRPCUnaryClientInterceptor(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		md = make(metadata.MD, len(tracingKeys))
	}

	for _, key := range tracingKeys {
		mdKey := strings.ToLower(key)

		if _, ok := md[mdKey]; ok {
			delete(md, mdKey)
		}

		val := ctx.Value(key)
		if val == nil {
			md[mdKey] = []string{""}
		} else {
			md[mdKey] = []string{val.(string)}
		}
	}

	ctx = metadata.NewOutgoingContext(ctx, md)

	return invoker(ctx, method, req, reply, cc, opts...)
}
