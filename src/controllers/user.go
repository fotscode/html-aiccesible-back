package controllers

import (
	"html-aiccesible/httputil"
	"html-aiccesible/models"
	"os"
	"time"

	"html-aiccesible/repositories"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func (b *Controller) CreateUser(c *gin.Context) {
	body := c.MustGet(gin.BindKey).(*models.CreateUserBody)
	user, err := repositories.UserRepo().CreateUser(body)
	if err != nil {
		httputil.InternalServerError[string](c, err.Error())
		return
	}
	httputil.Created[*models.User](c, user)
}

func (b *Controller) UpdateUser(c *gin.Context) {
	body := c.MustGet(gin.BindKey).(*models.UpdateUserBody)
	user, err := repositories.UserRepo().UpdateUser(body)
	if err != nil {
		httputil.InternalServerError[string](c, err.Error())
		return
	}
	httputil.OK[*models.User](c, user)
}

func (b *Controller) GetUser(c *gin.Context) {
	getOpt := c.MustGet("getOpt").(*models.GetOptions)
	user, err := repositories.UserRepo().GetUser(getOpt.Id)
	if err != nil {
		httputil.NotFound(c, err.Error())
		return
	}
	httputil.OK[*models.User](c, user)
}

func (b *Controller) DeleteUser(c *gin.Context) {
	getOpt := c.MustGet("getOpt").(*models.GetOptions)
	err := repositories.UserRepo().DeleteUser(getOpt.Id)
	if err != nil {
		httputil.InternalServerError[string](c, err.Error())
		return
	}
	httputil.NoContent[string](c, "Deleted user successfully")
}

func (b *Controller) ListUsers(c *gin.Context) {
	lo := c.MustGet("lo").(*models.ListOptions)
	users, err := repositories.UserRepo().ListUsers(lo.Page, lo.Size)
	if err != nil {
		httputil.InternalServerError[string](c, err.Error())
		return
	}
	httputil.OK[*[]models.User](c, &users)
}

func (b *Controller) Login(c *gin.Context) {
	body := c.MustGet(gin.BindKey).(*models.LoginUserBody)
	user, err := repositories.UserRepo().GetUserByUsername(body.Username)
	if err != nil {
		httputil.InternalServerError[string](c, err.Error())
		return
	}
	if user == nil {
		httputil.Unauthorized[string](c, "Invalid username or password")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		httputil.Unauthorized[string](c, "Invalid username or password")
		return
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}).SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		httputil.InternalServerError[string](c, err.Error())
		return
	}
	httputil.OK[*models.LoginResponse](c, &models.LoginResponse{Token: token})
}
