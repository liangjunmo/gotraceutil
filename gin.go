package gotraceutil

import (
	"context"

	"github.com/gin-gonic/gin"
)

func GinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		for _, key := range tracingKeys {
			ctx = context.WithValue(ctx, key, c.GetHeader(key))
		}

		if tracingID := ctx.Value(tracingKeys[0]); tracingID == "" {
			ctx = Trace(ctx)
		}

		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
