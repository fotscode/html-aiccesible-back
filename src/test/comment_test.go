package test

import (
	"fmt"
	"html-aiccesible/httputil"
	"html-aiccesible/models"
	"html-aiccesible/routes"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func createComment(t *testing.T, r *gin.Engine, postID uint, token string) *models.Comment {
	var res httputil.HTTPResponse[*models.Comment]
	w := createRequest(t, r, http.MethodPost, "/api/comment/add", models.CreateCommentBody{
		PostID:  postID,
		Title:   generateRandomString(10),
		Content: generateRandomString(10),
	}, &res, token)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected %d but got %d", http.StatusCreated, res.Code)
	}
	return res.Data
}

func TestGetComment(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := routes.SetUpRouter()

	_, token := login(t, r, false)
	post := createPost(t, r, token)
	comment := createComment(t, r, post.ID, token)

	tests := []TestBody[string]{
		{
			Name:         "Get comment by id successfully",
			Path:         fmt.Sprintf("/%d", comment.ID),
			ExpectedCode: http.StatusOK,
			RespContains: comment.Title,
		},
		{
			Name:         "Get comment by id not found",
			Path:         "/262144",
			ExpectedCode: http.StatusNotFound,
			RespContains: "",
		},
		{
			Name:         "Get comment by id invalid id",
			Path:         "/invalid",
			ExpectedCode: http.StatusBadRequest,
			RespContains: "Invalid ID",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var res httputil.HTTPResponse[interface{}]
			w := createRequest(t, r, http.MethodGet, "/api/comment/get"+test.Path, test.Body, &res, test.Token)
			doAsserts(t, w, res, test)
		})
	}
}

func TestListComment(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := routes.SetUpRouter()

	tests := []TestBody[string]{}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var res httputil.HTTPResponse[interface{}]
			w := createRequest(t, r, http.MethodGet, "/api/comment/list"+test.Path, test.Body, &res, test.Token)
			doAsserts(t, w, res, test)
		})
	}
}

func TestCreateComment(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := routes.SetUpRouter()

	tests := []TestBody[models.CreateCommentBody]{}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var res httputil.HTTPResponse[interface{}]
			w := createRequest(t, r, http.MethodPost, "/api/comment/add", test.Body, &res, test.Token)
			doAsserts(t, w, res, test)
		})
	}
}

func TestUpdateComment(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := routes.SetUpRouter()

	tests := []TestBody[models.UpdateCommentBody]{}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var res httputil.HTTPResponse[interface{}]
			w := createRequest(t, r, http.MethodPut, "/api/comment/update", test.Body, &res, test.Token)
			doAsserts(t, w, res, test)
		})
	}
}

func TestDeleteComment(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := routes.SetUpRouter()

	tests := []TestBody[string]{}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var res httputil.HTTPResponse[interface{}]
			w := createRequest(t, r, http.MethodDelete, "/api/comment/delete"+test.Path, test.Body, &res, test.Token)
			doAsserts(t, w, res, test)
		})
	}
}
