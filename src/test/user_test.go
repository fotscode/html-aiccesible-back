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
			w := createRequest(t, r, http.MethodPost, "/api/user/add", test.Body, &res, test.Token)
			doAsserts(t, w, res, test)

		})
	}
}

func TestUpdateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := routes.SetUpRouter()

	randomString := generateRandomString(10)
	var res httputil.HTTPResponse[*models.User]
	w := createRequest(t, r, http.MethodPost, "/api/user/add", models.CreateUserBody{
		Username: "userUpdate" + randomString,
		Password: "password",
	}, &res, "")

	if w.Code != http.StatusCreated {
		t.Errorf("Expected %d but got %d", http.StatusCreated, res.Code)
	}

	var resExists httputil.HTTPResponse[*models.User]
	createRequest(t, r, http.MethodPost, "/api/user/add", models.CreateUserBody{
		Username: "exists" + randomString,
		Password: "password",
	}, &resExists, "")

	var loginRes httputil.HTTPResponse[*models.LoginResponse]
	w = createRequest(t, r, http.MethodPost, "/api/user/login", models.LoginUserBody{
		Username: "userUpdate" + randomString,
		Password: "password",
	}, &loginRes, "")

	if w.Code != http.StatusOK {
		t.Errorf("Expected %d but got %d", http.StatusOK, w.Code)
	}

	token := loginRes.Data.Token

	updateRandomString := generateRandomString(10)
	tests := []TestBody[models.UpdateUserBody]{
		{
			Name: "Update user successfully",
			Body: models.UpdateUserBody{
				ID:       res.Data.ID,
				Username: "userUpdate" + updateRandomString,
				Password: "password",
			},
			ExpectedCode: http.StatusOK,
			RespContains: updateRandomString,
			Token:        token,
		},
		{
			Name: "Update user with empty username",
			Body: models.UpdateUserBody{
				ID:       res.Data.ID,
				Password: "password",
			},
			ExpectedCode: http.StatusBadRequest,
			RespContains: "Username is required",
			Token:        token,
		},
		{
			Name: "Update user with empty password",
			Body: models.UpdateUserBody{
				ID:       res.Data.ID,
				Username: "userUpdate" + updateRandomString,
			},
			ExpectedCode: http.StatusBadRequest,
			RespContains: "required",
			Token:        token,
		},
		{
			Name:         "Update user with empty username and password, check user",
			Body:         models.UpdateUserBody{},
			ExpectedCode: http.StatusBadRequest,
			RespContains: "Username is required",
			Token:        token,
		},
		{
			Name:         "Update user with empty username and password, check password",
			Body:         models.UpdateUserBody{},
			ExpectedCode: http.StatusBadRequest,
			RespContains: "Password is required",
			Token:        token,
		},
		{
			Name: "Update user with username less than 4 characters",
			Body: models.UpdateUserBody{
				ID:       res.Data.ID,
				Username: "tes",
				Password: "password",
			},
			ExpectedCode: http.StatusBadRequest,
			RespContains: "Username must be longer than 4",
			Token:        token,
		},
		{
			Name: "Update user with password less than 8 characters",
			Body: models.UpdateUserBody{
				ID:       res.Data.ID,
				Username: "userUpdate" + updateRandomString,
				Password: "passwor",
			},
			ExpectedCode: http.StatusBadRequest,
			RespContains: "Password must be longer than 8",
			Token:        token,
		},
		{
			Name: "Update user with username longer than 20 characters",
			Body: models.UpdateUserBody{
				ID:       res.Data.ID,
				Username: generateRandomString(21),
				Password: "password",
			},
			ExpectedCode: http.StatusBadRequest,
			RespContains: "Username cannot be longer than 20",
			Token:        token,
		},
		{
			Name: "Update user with password longer than 20 characters",
			Body: models.UpdateUserBody{
				ID:       res.Data.ID,
				Username: "userUpdate" + updateRandomString,
				Password: generateRandomString(21),
			},
			ExpectedCode: http.StatusBadRequest,
			RespContains: "Password cannot be longer than 20",
			Token:        token,
		},
		{
			Name: "Update user with username that already exists",
			Body: models.UpdateUserBody{
				ID:       res.Data.ID,
				Username: "exists" + randomString,
				Password: "password",
			},
			ExpectedCode: http.StatusInternalServerError,
			RespContains: "Duplicate entry",
			Token:        token,
		},
        {
            Name: "Update non logged in user",
            Body: models.UpdateUserBody{
                ID:       resExists.Data.ID,
                Username: "exists" + updateRandomString,
                Password: "password",
            },
            ExpectedCode: http.StatusUnauthorized,
            RespContains: "You are not authorized",
            Token:        token,
        },
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var res httputil.HTTPResponse[interface{}]
			w := createRequest(t, r, http.MethodPut, "/api/user/update", test.Body, &res, test.Token)
			doAsserts(t, w, res, test)

		})
	}

}
