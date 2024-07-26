package httputil

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HTTPResponse[T any] struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   T      `json:"data"`
}

func newHTTPResponse[T any](code int, data T) HTTPResponse[T] {
	return HTTPResponse[T]{
		Code:   code,
		Status: http.StatusText(code),
		Data:   data,
	}
}

// send using gin
func OK[T any](c *gin.Context, data T) {
	c.JSON(http.StatusOK, newHTTPResponse(http.StatusOK, data))
}

func Created[T any](c *gin.Context, data T) {
	c.JSON(http.StatusCreated, newHTTPResponse(http.StatusOK, data))
}

func BadRequest[T any](c *gin.Context, data T) {
	c.AbortWithStatusJSON(http.StatusBadRequest, newHTTPResponse(http.StatusBadRequest, data))
}

func InternalServerError[T any](c *gin.Context, data T) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, newHTTPResponse(http.StatusInternalServerError, data))
}

func NoContent[T any](c *gin.Context, data T) {
	c.AbortWithStatusJSON(http.StatusNoContent, newHTTPResponse(http.StatusNoContent, data))
}

func Unauthorized[T any](c *gin.Context, data T) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, newHTTPResponse(http.StatusUnauthorized, data))
}
