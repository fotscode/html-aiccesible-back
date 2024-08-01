package test

import (
	"bytes"
	"encoding/json"
	"html-aiccesible/constants"
	"html-aiccesible/httputil"
	"html-aiccesible/models"
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
	Path         string
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

func createUser(t *testing.T, r *gin.Engine, username, password string) *models.User {
	var res httputil.HTTPResponse[*models.User]
	w := createRequest(t, r, http.MethodPost, "/api/user/add", models.CreateUserBody{
		Username: username,
		Password: password,
	}, &res, "")

	if w.Code != http.StatusCreated {
		t.Errorf("Expected %d but got %d", http.StatusCreated, res.Code)
	}

	return res.Data
}

func login(t *testing.T, r *gin.Engine, isAdmin bool) (*models.User, string) {
	randomString := generateRandomString(20)
	username := constants.ADMIN_USERNAME
	password := constants.ADMIN_PASSWORD
	var user *models.User
	if !isAdmin {
		username = randomString
		password = "password"
		user = createUser(t, r, username, password)
	} else {
		models.CreateDefaultUser()
		user = &models.User{
			Username: username,
			Password: password,
		}

	}
	var loginRes httputil.HTTPResponse[models.LoginResponse]
	w := createRequest(t, r, http.MethodPost, "/api/user/login", models.LoginUserBody{
		Username: username,
		Password: password,
	}, &loginRes, "")

	if w.Code != http.StatusOK {
		t.Errorf("Expected %d but got %d", http.StatusOK, w.Code)
	}
	return user, loginRes.Data.Token
}
