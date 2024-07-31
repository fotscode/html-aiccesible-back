package test

import (
	"bytes"
	"encoding/json"
	"html-aiccesible/httputil"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type TestBody[T any] struct {
	Name         string
	Body         T
	ExpectedCode int
	RespContains string
	Token        string
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func generateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func createRequest[T any, R any](t *testing.T, r *gin.Engine, method, path string, body T, expected *httputil.HTTPResponse[R], token string) *httptest.ResponseRecorder {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)
	req, err := http.NewRequest(method, path, buf)
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	json.Unmarshal(w.Body.Bytes(), &expected)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}
	return w
}

func doAsserts[T any, K any](t *testing.T, w *httptest.ResponseRecorder, expected httputil.HTTPResponse[T], test TestBody[K]) {
	assert.Equal(t, w.Code, expected.Code)
	assert.Equal(t, test.ExpectedCode, expected.Code)
	if test.RespContains != "" {
		buf, err := json.Marshal(expected.Data)
		if err != nil {
			t.Fatalf("could not marshal expected data: %v", err)
		}
		bufStr := string(buf)
		assert.Contains(t, bufStr, test.RespContains)
	}
}
