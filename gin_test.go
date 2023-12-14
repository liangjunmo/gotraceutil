package gotraceutil_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/levigross/grequests"
	"github.com/stretchr/testify/assert"

	"github.com/liangjunmo/gotraceutil"
)

func TestGinMiddleware(t *testing.T) {
	tracingIDKey := "TracingID"
	tracingID := "tracingID"

	clientIDKey := "ClientID"
	clientID := "clientID"

	gotraceutil.SetTracingKeys([]string{tracingIDKey, clientIDKey})

	gotraceutil.SetTracingIDGenerator(func() string {
		return tracingID
	})

	router := gin.Default()

	router.Use(gotraceutil.GinMiddleware())

	router.GET("/", func(c *gin.Context) {
		ctx := c.Request.Context()
		labels := gotraceutil.Parse(ctx)
		assert.Equal(t, tracingID, labels[tracingIDKey])
		assert.Equal(t, clientID, labels[clientIDKey])
	})

	server := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	go func() {
		err := server.ListenAndServe()
		if !errors.Is(err, http.ErrServerClosed) {
			t.Error(err)
			return
		}
	}()

	_, err := grequests.Get("http://127.0.0.1:8000/", &grequests.RequestOptions{
		Headers: map[string]string{
			clientIDKey: clientID,
		},
	})
	assert.Nil(t, err)

	err = server.Shutdown(context.Background())
	assert.Nil(t, err)
}
