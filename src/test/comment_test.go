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

	_, token := login(t, r, false)
	post := createPost(t, r, token)
	comment1 := createComment(t, r, post.ID, token)
	comment2 := createComment(t, r, post.ID, token)

	tests := []TestBody[string]{
		{
			Name:         "List comments successfully pt1",
			Path:         fmt.Sprintf("/%d", post.ID),
			ExpectedCode: http.StatusOK,
			RespContains: comment1.Title,
		},
		{
			Name:         "List comments successfully pt2",
			Path:         fmt.Sprintf("/%d", post.ID),
			ExpectedCode: http.StatusOK,
			RespContains: comment2.Title,
		},
		{
			Name:         "List comments invalid id",
			Path:         "/invalid",
			ExpectedCode: http.StatusBadRequest,
			RespContains: "Invalid ID",
		},
		{
			Name:         "List comments id not found",
			Path:         "/262144",
			ExpectedCode: http.StatusNotFound,
			RespContains: "",
		},
	}

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

	_, token := login(t, r, false)
	post := createPost(t, r, token)

	title := generateRandomString(10)
	content := generateRandomString(10)
	tests := []TestBody[models.CreateCommentBody]{
		{
			Name: "Create comment successfully",
			Body: models.CreateCommentBody{
				PostID:  post.ID,
				Title:   title,
				Content: content,
			},
			Token:        token,
			ExpectedCode: http.StatusCreated,
			RespContains: title,
		},
		{
			Name: "Create comment invalid post id",
			Body: models.CreateCommentBody{
				PostID:  262144,
				Title:   title,
				Content: content,
			},
			Token:        token,
			ExpectedCode: http.StatusInternalServerError,
			RespContains: "",
		},
		{
			Name: "Create comment no title",
			Body: models.CreateCommentBody{
				PostID:  post.ID,
				Content: content,
			},
			Token:        token,
			ExpectedCode: http.StatusBadRequest,
			RespContains: "Title is required",
		},
		{
			Name: "Create comment no content",
			Body: models.CreateCommentBody{
				PostID: post.ID,
				Title:  title,
			},
			Token:        token,
			ExpectedCode: http.StatusBadRequest,
			RespContains: "Content is required",
		},
		{
			Name: "Create comment no title and content",
			Body: models.CreateCommentBody{
				PostID: post.ID,
			},
			Token:        token,
			ExpectedCode: http.StatusBadRequest,
			RespContains: "\"Content\":\"Content is required\",\"Title\":\"Title is required\"",
		},
		{
			Name: "Create comment no token",
			Body: models.CreateCommentBody{
				PostID:  post.ID,
				Title:   title,
				Content: content,
			},
			ExpectedCode: http.StatusUnauthorized,
			RespContains: "No token provided",
		},
		{
			Name: "Create comment with title less than 5",
			Body: models.CreateCommentBody{
				PostID:  post.ID,
				Title:   "1234",
				Content: content,
			},
			Token:        token,
			ExpectedCode: http.StatusBadRequest,
			RespContains: "Title must be longer than 5",
		},
		{
			Name: "Create comment with content less than 5",
			Body: models.CreateCommentBody{
				PostID:  post.ID,
				Title:   title,
				Content: "1234",
			},
			Token:        token,
			ExpectedCode: http.StatusBadRequest,
			RespContains: "Content must be longer than 5",
		},
		{
			Name: "Create comment with title more than 100",
			Body: models.CreateCommentBody{
				PostID:  post.ID,
				Title:   generateRandomString(101),
				Content: content,
			},
			Token:        token,
			ExpectedCode: http.StatusBadRequest,
			RespContains: "Title cannot be longer than 100",
		},
		{
			Name: "Create comment with content more than 1000",
			Body: models.CreateCommentBody{
				PostID:  post.ID,
				Title:   title,
				Content: generateRandomString(1001),
			},
			Token:        token,
			ExpectedCode: http.StatusBadRequest,
			RespContains: "Content cannot be longer than 1000",
		},
	}

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
