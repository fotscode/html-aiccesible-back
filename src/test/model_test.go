package test

import (
	"html-aiccesible/httputil"
	routes "html-aiccesible/routes"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestListModels(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := routes.SetUpRouter()
	tests := []TestBody[string]{
		{
			Name:         "List models successfully",
			ExpectedCode: http.StatusOK,
			RespContains: "[",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var res httputil.HTTPResponse[[]string]
			w := createRequest(t, r, http.MethodGet, "/api/models/list", test.Body, &res, test.Token)
			doAsserts(t, w, res, test)
		})
	}
}
