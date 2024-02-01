package gotraceutil

import (
	"context"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func GRPCUnaryServerInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	md, exist := metadata.FromIncomingContext(ctx)

	for _, key := range tracingKeys {
		if !exist {
			ctx = context.WithValue(ctx, key, "")
			continue
		}

		mdKey := strings.ToLower(key)

		vals, ok := md[mdKey]
		if !ok {
			ctx = context.WithValue(ctx, key, "")
		} else if len(vals) == 0 {
			ctx = context.WithValue(ctx, key, "")
		} else {
			ctx = context.WithValue(ctx, key, vals[0])
		}
	}

	if ctx.Value(tracingIDKey) == "" {
		ctx = Trace(ctx)
	}

	return handler(ctx, req)
}

func GRPCUnaryClientInterceptor(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	md, exist := metadata.FromOutgoingContext(ctx)
	if !exist {
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
