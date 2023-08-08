package gotraceutil

import (
	"context"

	"github.com/gin-gonic/gin"
)

func GinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		for _, key := range traceKeys {
			ctx = context.WithValue(ctx, key, c.GetHeader(key))
		}

		if traceId := ctx.Value(traceIdKey); traceId == "" {
			ctx = Trace(ctx)
		}

		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
