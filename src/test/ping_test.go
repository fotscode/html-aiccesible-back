package test

import (
	"encoding/json"
	"html-aiccesible/httputil"
	routes "html-aiccesible/routes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	r := routes.SetUpRouter()
	req, _ := http.NewRequest("GET", "/api/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	var expected httputil.HTTPResponse[string]
	json.Unmarshal(w.Body.Bytes(), &expected)
	assert.Equal(t, w.Code, expected.Code)
	assert.Equal(t, "pong", expected.Data)
}

func TestProtectedPingRouteTokenless(t *testing.T) {
	r := routes.SetUpRouter()
	req, _ := http.NewRequest("GET", "/api/protected", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	var expected httputil.HTTPResponse[string]
	json.Unmarshal(w.Body.Bytes(), &expected)
	assert.Equal(t, w.Code, http.StatusUnauthorized)
	assert.Equal(t, "No token provided", expected.Data)
}
