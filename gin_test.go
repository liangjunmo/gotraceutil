package gotraceutil

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/levigross/grequests"
	"github.com/stretchr/testify/require"
)

func TestGinMiddleware(t *testing.T) {
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

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.Use(GinMiddleware())

	router.GET("/", func(c *gin.Context) {
		ctx := c.Request.Context()
		require.Equal(t, tracingIDValue, ctx.Value(tracingIDKey))
		require.Equal(t, clientIDValue, ctx.Value(clientIDKey))
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
			clientIDKey: clientIDValue,
		},
	})
	require.Nil(t, err)

	err = server.Shutdown(context.Background())
	require.Nil(t, err)
}
