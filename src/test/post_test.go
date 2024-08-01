package test

import (
	"fmt"
	"html-aiccesible/httputil"
	"html-aiccesible/models"
	"html-aiccesible/repositories"
	routes "html-aiccesible/routes"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func createPost(t *testing.T, r *gin.Engine, token string) *models.Post {
	var res httputil.HTTPResponse[*models.Post]
	w := createRequest(t, r, http.MethodPost, "/api/post/add", models.CreatePostBody{
		Title:       generateRandomString(10),
		Description: generateRandomString(10),
		Before:      generateRandomString(10),
		After:       generateRandomString(10),
	}, &res, token)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected %d but got %d", http.StatusCreated, res.Code)
	}
	return res.Data
}

func TestGetPost(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := routes.SetUpRouter()

	_, token := login(t, r, false)
	post := createPost(t, r, token)

	tests := []TestBody[string]{
		{
			Name:         "Get post successfully",
			Path:         fmt.Sprintf("/%d", post.ID),
			ExpectedCode: http.StatusOK,
			RespContains: post.Title,
		},
		{
			Name:         "Get post with invalid ID",
			Path:         "/invalid",
			ExpectedCode: http.StatusBadRequest,
			RespContains: "Invalid ID",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var res httputil.HTTPResponse[interface{}]
			w := createRequest(t, r, http.MethodGet, "/api/post/get"+test.Path, test.Body, &res, test.Token)
			doAsserts(t, w, res, test)
		})
	}
}

func TestListPost(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := routes.SetUpRouter()

	tests := []TestBody[string]{
		{
			Name:         "List post successfully",
			ExpectedCode: http.StatusOK,
			RespContains: "[",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var res httputil.HTTPResponse[interface{}]
			w := createRequest(t, r, http.MethodGet, "/api/post/list", test.Body, &res, test.Token)
			doAsserts(t, w, res, test)
		})
	}
}

func TestAddPost(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := routes.SetUpRouter()

	_, token := login(t, r, false)
	title := generateRandomString(10)
	description := generateRandomString(10)
	before := generateRandomString(10)
	after := generateRandomString(10)

	tests := []TestBody[models.CreatePostBody]{
		{
			Name: "Add post successfully",
			Body: models.CreatePostBody{
				Title:       title,
				Description: description,
				Before:      before,
				After:       after,
			},
			Token:        token,
			ExpectedCode: http.StatusCreated,
			RespContains: title,
		},
		{
			Name: "Add post with no token",
			Body: models.CreatePostBody{
				Title:       title,
				Description: description,
				Before:      before,
				After:       after,
			},
			ExpectedCode: http.StatusUnauthorized,
			RespContains: "No token provided",
		},
		{
			Name: "Add post with no title",
			Body: models.CreatePostBody{
				Description: description,
				Before:      before,
				After:       after,
			},
			Token:        token,
			ExpectedCode: http.StatusBadRequest,
			RespContains: "Title is required",
		},
		{
			Name: "Add post with no description",
			Body: models.CreatePostBody{
				Title:  title,
				Before: before,
				After:  after,
			},
			Token:        token,
			ExpectedCode: http.StatusBadRequest,
			RespContains: "Description is required",
		},
		{
			Name: "Add post with no before",
			Body: models.CreatePostBody{
				Title:       title,
				Description: description,
				After:       after,
			},
			Token:        token,
			ExpectedCode: http.StatusBadRequest,
			RespContains: "Before is required",
		},
		{
			Name: "Add post with no after",
			Body: models.CreatePostBody{
				Title:       title,
				Description: description,
				Before:      before,
			},
			Token:        token,
			ExpectedCode: http.StatusBadRequest,
			RespContains: "After is required",
		},
		{
			Name:         "Add empty post",
			Body:         models.CreatePostBody{},
			Token:        token,
			ExpectedCode: http.StatusBadRequest,
			RespContains: "{\"After\":\"After is required\",\"Before\":\"Before is required\",\"Description\":\"Description is required\",\"Title\":\"Title is required\"}",
		},
		{
			Name: "Add post with title length less than 4",
			Body: models.CreatePostBody{
				Title:       "abc",
				Description: description,
				Before:      before,
				After:       after,
			},
			Token:        token,
			ExpectedCode: http.StatusBadRequest,
			RespContains: "Title must be longer than 4",
		},
		{
			Name: "Add post with description length less than 4",
			Body: models.CreatePostBody{
				Title:       title,
				Description: "abc",
				Before:      before,
				After:       after,
			},
			Token:        token,
			ExpectedCode: http.StatusBadRequest,
			RespContains: "Description must be longer than 4",
		},
		{
			Name: "Add post with before length less than 4",
			Body: models.CreatePostBody{
				Title:       title,
				Description: description,
				Before:      "abc",
				After:       after,
			},
			Token:        token,
			ExpectedCode: http.StatusBadRequest,
			RespContains: "Before must be longer than 4",
		},
		{
			Name: "Add post with after length less than 4",
			Body: models.CreatePostBody{
				Title:       title,
				Description: description,
				Before:      before,
				After:       "abc",
			},
			Token:        token,
			ExpectedCode: http.StatusBadRequest,
			RespContains: "After must be longer than 4",
		},
		{
			Name: "Add post with title length more than 100",
			Body: models.CreatePostBody{
				Title:       generateRandomString(101),
				Description: description,
				Before:      before,
				After:       after,
			},
			Token:        token,
			ExpectedCode: http.StatusBadRequest,
			RespContains: "Title cannot be longer than 100",
		},
		{
			Name: "Add post with description length more than 100",
			Body: models.CreatePostBody{
				Title:       title,
				Description: generateRandomString(101),
				Before:      before,
				After:       after,
			},
			Token:        token,
			ExpectedCode: http.StatusBadRequest,
			RespContains: "Description cannot be longer than 100",
		},
		{
			Name: "Add post with before length more than 8192",
			Body: models.CreatePostBody{
				Title:       title,
				Description: description,
				Before:      generateRandomString(8193),
				After:       after,
			},
			Token:        token,
			ExpectedCode: http.StatusBadRequest,
			RespContains: "Before cannot be longer than 8192",
		},
		{
			Name: "Add post with after length more than 8192",
			Body: models.CreatePostBody{
				Title:       title,
				Description: description,
				Before:      before,
				After:       generateRandomString(8193),
			},
			Token:        token,
			ExpectedCode: http.StatusBadRequest,
			RespContains: "After cannot be longer than 8192",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var res httputil.HTTPResponse[interface{}]
			w := createRequest(t, r, http.MethodPost, "/api/post/add", test.Body, &res, test.Token)
			doAsserts(t, w, res, test)
		})
	}
}

func TestUpdatePost(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := routes.SetUpRouter()

	_, token := login(t, r, false)
	title := generateRandomString(10)
	description := generateRandomString(10)
	before := generateRandomString(10)
	after := generateRandomString(10)
	post := createPost(t, r, token)

	_, otherToken := login(t, r, false)

	tests := []TestBody[models.UpdatePostBody]{
		{
			Name: "Update post successfully",
			Body: models.UpdatePostBody{
				ID:          post.ID,
				Title:       title,
				Description: description,
				Before:      before,
				After:       after,
			},
			Token:        token,
			ExpectedCode: http.StatusOK,
			RespContains: title,
		},
		{
			Name: "Update post with other user token",
			Body: models.UpdatePostBody{
				ID:          post.ID,
				Title:       title,
				Description: description,
				Before:      before,
				After:       after,
			},
			Token:        otherToken,
			ExpectedCode: http.StatusInternalServerError,
			RespContains: "user is not the owner of the post",
		},
		{
			Name: "Update post with no token",
			Body: models.UpdatePostBody{
				ID:          post.ID,
				Title:       title,
				Description: description,
				Before:      before,
				After:       after,
			},
			ExpectedCode: http.StatusUnauthorized,
			RespContains: "No token provided",
		},
		{
			Name: "Update post with no ID",
			Body: models.UpdatePostBody{
				Title:       title,
				Description: description,
				Before:      before,
				After:       after,
			},
			Token:        token,
			ExpectedCode: http.StatusBadRequest,
			RespContains: "ID is required",
		},
		{
			Name: "Update post with no title",
			Body: models.UpdatePostBody{
				ID:          post.ID,
				Description: description,
				Before:      before,
				After:       after,
			},
			Token:        token,
			ExpectedCode: http.StatusBadRequest,
			RespContains: "Title is required",
		},
		{
			Name: "Update post with no description",
			Body: models.UpdatePostBody{
				ID:     post.ID,
				Title:  title,
				Before: before,
				After:  after,
			},
			Token:        token,
			ExpectedCode: http.StatusBadRequest,
			RespContains: "Description is required",
		},
		{
			Name: "Update post with no before",
			Body: models.UpdatePostBody{
				ID:          post.ID,
				Title:       title,
				Description: description,
				After:       after,
			},
			Token:        token,
			ExpectedCode: http.StatusBadRequest,
			RespContains: "Before is required",
		},
		{
			Name: "Update post with no after",
			Body: models.UpdatePostBody{
				ID:          post.ID,
				Title:       title,
				Description: description,
				Before:      before,
			},
			Token:        token,
			ExpectedCode: http.StatusBadRequest,
			RespContains: "After is required",
		},
		{
			Name:         "Update empty post",
			Body:         models.UpdatePostBody{},
			Token:        token,
			ExpectedCode: http.StatusBadRequest,
			RespContains: "{\"After\":\"After is required\",\"Before\":\"Before is required\",\"Description\":\"Description is required\",\"ID\":\"ID is required\",\"Title\":\"Title is required\"}",
		},
		{
			Name: "Update post with title length less than 4",
			Body: models.UpdatePostBody{
				ID:          post.ID,
				Title:       "abc",
				Description: description,
				Before:      before,
				After:       after,
			},
			Token:        token,
			ExpectedCode: http.StatusBadRequest,
			RespContains: "Title must be longer than 4",
		},
		{
			Name: "Update post with description length less than 4",
			Body: models.UpdatePostBody{
				ID:          post.ID,
				Title:       title,
				Description: "abc",
				Before:      before,
				After:       after,
			},
			Token:        token,
			ExpectedCode: http.StatusBadRequest,
			RespContains: "Description must be longer than 4",
		},
		{
			Name: "Update post with before length less than 4",
			Body: models.UpdatePostBody{
				ID:          post.ID,
				Title:       title,
				Description: description,
				Before:      "abc",
				After:       after,
			},
			Token:        token,
			ExpectedCode: http.StatusBadRequest,
			RespContains: "Before must be longer than 4",
		},
		{
			Name: "Update post with after length less than 4",
			Body: models.UpdatePostBody{
				ID:          post.ID,
				Title:       title,
				Description: description,
				Before:      before,
				After:       "abc",
			},
			Token:        token,
			ExpectedCode: http.StatusBadRequest,
			RespContains: "After must be longer than 4",
		},
		{
			Name: "Update post with title length more than 100",
			Body: models.UpdatePostBody{
				ID:          post.ID,
				Title:       generateRandomString(101),
				Description: description,
				Before:      before,
				After:       after,
			},
			Token:        token,
			ExpectedCode: http.StatusBadRequest,
			RespContains: "Title cannot be longer than 100",
		},
		{
			Name: "Update post with description length more than 100",
			Body: models.UpdatePostBody{
				ID:          post.ID,
				Title:       title,
				Description: generateRandomString(101),
				Before:      before,
				After:       after,
			},
			Token:        token,
			ExpectedCode: http.StatusBadRequest,
			RespContains: "Description cannot be longer than 100",
		},
		{
			Name: "Update post with before length more than 8192",
			Body: models.UpdatePostBody{
				ID:          post.ID,
				Title:       title,
				Description: description,
				Before:      generateRandomString(8193),
				After:       after,
			},
			Token:        token,
			ExpectedCode: http.StatusBadRequest,
			RespContains: "Before cannot be longer than 8192",
		},
		{
			Name: "Update post with after length more than 8192",
			Body: models.UpdatePostBody{
				ID:          post.ID,
				Title:       title,
				Description: description,
				Before:      before,
				After:       generateRandomString(8193),
			},
			Token:        token,
			ExpectedCode: http.StatusBadRequest,
			RespContains: "After cannot be longer than 8192",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var res httputil.HTTPResponse[interface{}]
			w := createRequest(t, r, http.MethodPut, "/api/post/update", test.Body, &res, test.Token)
			doAsserts(t, w, res, test)
		})
	}
}

func TestDeletePost(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := routes.SetUpRouter()

	_, token := login(t, r, false)
	post := createPost(t, r, token)
	postOther := createPost(t, r, token)
	_, otherToken := login(t, r, false)

	tests := []TestBody[string]{
		{
			Name:         "Delete post successfully",
			Path:         fmt.Sprintf("/%d", post.ID),
			Token:        token,
			ExpectedCode: http.StatusOK,
			RespContains: "Deleted post",
		},
		{
			Name:         "Delete post with invalid ID",
			Path:         "/invalid",
			Token:        token,
			ExpectedCode: http.StatusBadRequest,
			RespContains: "Invalid ID",
		},
		{
			Name:         "Delete post with no token",
			Path:         fmt.Sprintf("/%d", post.ID),
			ExpectedCode: http.StatusUnauthorized,
			RespContains: "No token provided",
		},
		{
			Name:         "Delete post with other user token",
			Path:         fmt.Sprintf("/%d", postOther.ID),
			Token:        otherToken,
			ExpectedCode: http.StatusInternalServerError,
			RespContains: "user is not the owner of the post",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var res httputil.HTTPResponse[interface{}]
			w := createRequest(t, r, http.MethodDelete, "/api/post/delete"+test.Path, test.Body, &res, test.Token)
			doAsserts(t, w, res, test)
		})
	}
}

func TestLikePost(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := routes.SetUpRouter()

	_, token := login(t, r, false)
	post := createPost(t, r, token)

	tests := []TestBody[string]{
		{
			Name:         "Like post successfully",
			Path:         fmt.Sprintf("/%d", post.ID),
			Token:        token,
			ExpectedCode: http.StatusOK,
			RespContains: "Toggle like post successfully",
		},
		{
			Name:         "Like post with invalid ID",
			Path:         "/invalid",
			Token:        token,
			ExpectedCode: http.StatusBadRequest,
			RespContains: "Invalid ID",
		},
		{
			Name:         "Like post with no token",
			Path:         fmt.Sprintf("/%d", post.ID),
			ExpectedCode: http.StatusUnauthorized,
			RespContains: "No token provided",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var res httputil.HTTPResponse[interface{}]
			w := createRequest(t, r, http.MethodPatch, "/api/post/like"+test.Path, test.Body, &res, test.Token)
			doAsserts(t, w, res, test)
		})
	}

	db := models.GetDB()
	searchPost, err := repositories.PostRepo(db).GetPost(int(post.ID))
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if len(searchPost.Likes) != 1 {
		t.Errorf("Expected 1 but got %d", len(searchPost.Likes))
	}
}

func TestGetLikePost(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := routes.SetUpRouter()

	_, token := login(t, r, false)
	post := createPost(t, r, token)

	tests := []TestBody[string]{
		{
			Name:         "Get like post successfully",
			Path:         fmt.Sprintf("/%d", post.ID),
			Token:        token,
			ExpectedCode: http.StatusOK,
			RespContains: "0",
		},
		{
			Name:         "Get like post with invalid ID",
			Path:         "/invalid",
			Token:        token,
			ExpectedCode: http.StatusBadRequest,
			RespContains: "Invalid ID",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var res httputil.HTTPResponse[interface{}]
			w := createRequest(t, r, http.MethodGet, "/api/post/likes"+test.Path, test.Body, &res, test.Token)
			doAsserts(t, w, res, test)
		})
	}

	likePost(t, r, token, post.ID)
	_, otherToken := login(t, r, false)
	likePost(t, r, otherToken, post.ID)

	test := TestBody[string]{
		Name:         "Get like post successfully with like",
		Path:         fmt.Sprintf("/%d", post.ID),
		Token:        token,
		ExpectedCode: http.StatusOK,
		RespContains: "2",
	}

	t.Run(test.Name, func(t *testing.T) {
		var res httputil.HTTPResponse[interface{}]
		w := createRequest(t, r, http.MethodGet, "/api/post/likes"+test.Path, test.Body, &res, test.Token)
		doAsserts(t, w, res, test)
	})

	likePost(t, r, otherToken, post.ID)

	test.RespContains = "1"

	t.Run(test.Name, func(t *testing.T) {
		var res httputil.HTTPResponse[interface{}]
		w := createRequest(t, r, http.MethodGet, "/api/post/likes"+test.Path, test.Body, &res, test.Token)
		doAsserts(t, w, res, test)
	})

}

func likePost(t *testing.T, r *gin.Engine, token string, postID uint) {
	var res httputil.HTTPResponse[interface{}]
	w := createRequest(t, r, http.MethodPatch, fmt.Sprintf("/api/post/like/%d", postID), "", &res, token)
	if w.Code != http.StatusOK {
		t.Errorf("Expected %d but got %d", http.StatusOK, res.Code)
	}
}
