package test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"html-aiccesible/httputil"
	routes "html-aiccesible/routes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingRoute(t *testing.T) {
	r := routes.SetUpRouter()
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	var expected httputil.HTTPOKResponse[string]
	json.Unmarshal(w.Body.Bytes(), &expected)
	assert.Equal(t, w.Code, expected.Code)
	assert.Equal(t, "pong", expected.Data)
}
