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

func TestGetConfig(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := routes.SetUpRouter()

	_, token := login(t, r, false)

	tests := []TestBody[string]{
		{
			Name:         "Get config successfully",
			Token:        token,
			ExpectedCode: http.StatusOK,
			RespContains: "UpdatedAt",
		},
		{
			Name:         "Get config with invalid token",
			ExpectedCode: http.StatusUnauthorized,
			RespContains: "No token provided",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var res httputil.HTTPResponse[interface{}]
			w := createRequest(t, r, http.MethodGet, "/api/config/get", test.Body, &res, test.Token)
			doAsserts(t, w, res, test)

		})
	}
}

func TestUpdateConfig(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := routes.SetUpRouter()
	user, token := login(t, r, false)
	db := models.GetDB()
	configRepo := repositories.ConfigRepo(db)
	config, err := configRepo.GetConfig(int(user.ID))
	if err != nil {
		t.Errorf("Error getting config: %v", err)
	}
	tests := []TestBody[models.UpdateConfigBody]{
		{
			Name: "Update config successfully",
			Body: models.UpdateConfigBody{
				ShowLikes:    !config.ShowLikes,
				ShowComments: !config.ShowComments,
				Theme:        config.Theme + "-test",
				Language:     config.Language + "-test",
				SizeTitle:    config.SizeTitle + 1,
				SizeText:     config.SizeText + 1,
			},
			ExpectedCode: http.StatusOK,
			RespContains: configStr(!config.ShowLikes, !config.ShowComments, config.Theme+"-test", config.Language+"-test", config.SizeTitle+1, config.SizeText+1),
			Token:        token,
		},
		{
			Name: "Update config with invalid token",
			Body: models.UpdateConfigBody{
				ShowLikes:    !config.ShowLikes,
				ShowComments: !config.ShowComments,
				Theme:        config.Theme + "-test",
				Language:     config.Language + "-test",
				SizeTitle:    config.SizeTitle + 1,
				SizeText:     config.SizeText + 1,
			},
			ExpectedCode: http.StatusUnauthorized,
			RespContains: "No token provided",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var res httputil.HTTPResponse[interface{}]
			w := createRequest(t, r, http.MethodPut, "/api/config/update", test.Body, &res, test.Token)
			doAsserts(t, w, res, test)

		})
	}
}

func configStr(showLikes, showComments bool, theme, language string, sizeTitle, sizeText int) string {
	return fmt.Sprintf("\"language\":\"%s\",\"show_comments\":%t,\"show_likes\":%t,\"size_text\":%d,\"size_title\":%d,\"theme\":\"%s\"", language, showComments, showLikes, sizeText, sizeTitle, theme)
}
