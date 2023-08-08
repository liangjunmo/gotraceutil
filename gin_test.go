package gotraceutil_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/levigross/grequests"
	"github.com/stretchr/testify/assert"

	"github.com/liangjunmo/gotraceutil"
)

func TestGinMiddleware(t *testing.T) {
	traceId := "trace-id-unique-string"
	clientId := "client-id-unique-string"

	gotraceutil.SetTraceIdGenerator(func() string {
		return traceId
	})

	gotraceutil.AppendTraceKeys([]string{"ClientId"})

	router := gin.Default()

	router.Use(gotraceutil.GinMiddleware())

	router.GET("/", func(c *gin.Context) {
		ctx := c.Request.Context()
		labels := gotraceutil.Parse(ctx)
		assert.Equal(t, traceId, labels["TraceId"])
		assert.Equal(t, clientId, labels["ClientId"])
	})

	server := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	go func() {
		err := server.ListenAndServe()
		if err != http.ErrServerClosed {
			t.Log(err)
		}
	}()

	_, err := grequests.Get("http://127.0.0.1:8000/", &grequests.RequestOptions{
		Headers: map[string]string{
			"ClientId": clientId,
		},
	})
	assert.Nil(t, err)

	err = server.Shutdown(context.Background())
	assert.Nil(t, err)
}
