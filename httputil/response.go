package httputil

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HTTPOKResponse[T any] struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   T      `json:"data"`
}

func newHTTPOKResponse[T any](code int, data T) HTTPOKResponse[T] {
	return HTTPOKResponse[T]{
		Code:   code,
		Status: http.StatusText(code),
		Data:   data,
	}
}

// send using gin
func OK[T any](c *gin.Context, data T) {
	c.JSON(http.StatusOK, newHTTPOKResponse(http.StatusOK, data))
}

func Created[T any](c *gin.Context, data T) {
	c.JSON(http.StatusCreated, newHTTPOKResponse(http.StatusOK, data))
}
