package test

import (
	"fmt"
	"html-aiccesible/httputil"
	"html-aiccesible/models"
	routes "html-aiccesible/routes"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCreateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := routes.SetUpRouter()

	randomString := generateRandomString(10)
	tests := []TestBody[models.CreateUserBody]{
		{
			Name: "Create user successfully",
			Body: models.CreateUserBody{
				Username: fmt.Sprintf("test-%s", randomString),
				Password: "password",
			},
			ExpectedCode: http.StatusCreated,
			RespContains: randomString,
		},
		{
			Name: "Create user with empty username",
			Body: models.CreateUserBody{
				Password: "password",
			},
			ExpectedCode: http.StatusBadRequest,
			RespContains: "Username is required",
		},
		{
			Name: "Create user with empty password",
			Body: models.CreateUserBody{
				Username: fmt.Sprintf("test-%s", generateRandomString(10)),
			},
			ExpectedCode: http.StatusBadRequest,
			RespContains: "required",
		},
		{
			Name:         "Create user with empty username and password, check user",
			Body:         models.CreateUserBody{},
			ExpectedCode: http.StatusBadRequest,
			RespContains: "Username is required",
		},
		{
			Name:         "Create user with empty username and password, check password",
			Body:         models.CreateUserBody{},
			ExpectedCode: http.StatusBadRequest,
			RespContains: "Password is required",
		},
		{
			Name: "Create user with username less than 4 characters",
			Body: models.CreateUserBody{
				Username: "tes",
				Password: "password",
			},
			ExpectedCode: http.StatusBadRequest,
			RespContains: "Username must be longer than 4",
		},
		{
			Name: "Create user with password less than 8 characters",
			Body: models.CreateUserBody{
				Username: fmt.Sprintf("test-%s", generateRandomString(10)),
				Password: "pass",
			},
			ExpectedCode: http.StatusBadRequest,
			RespContains: "Password must be longer than 8",
		},
		{
			Name: "Create user with username longer than 20 characters",
			Body: models.CreateUserBody{
				Username: generateRandomString(21),
				Password: "password",
			},
			ExpectedCode: http.StatusBadRequest,
			RespContains: "Username cannot be longer than 20",
		},
		{
			Name: "Create user with password longer than 20 characters",
			Body: models.CreateUserBody{
				Username: fmt.Sprintf("test-%s", generateRandomString(10)),
				Password: generateRandomString(21),
			},
			ExpectedCode: http.StatusBadRequest,
			RespContains: "Password cannot be longer than 20",
		},
		{
			Name: "Create user with username that already exists",
			Body: models.CreateUserBody{
				Username: "test",
				Password: "password",
			},
			ExpectedCode: http.StatusInternalServerError,
			RespContains: "Duplicate entry",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var res httputil.HTTPResponse[interface{}]
			w := createRequest(t, r, http.MethodPost, "/api/user/add", test.Body, &res)
			doAsserts(t, w, res, test)

		})
	}
}
