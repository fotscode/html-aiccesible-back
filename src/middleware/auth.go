package middleware

import (
	"fmt"
	"html-aiccesible/httputil"
	"html-aiccesible/models"
	"html-aiccesible/repositories"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			httputil.Unauthorized[string](c, "No token provided")
			return
		}

		if len(tokenString) < 7 || tokenString[:7] != "Bearer " {
			httputil.Unauthorized[string](c, "Invalid token")
			return
		}
		tokenString = tokenString[7:]

		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if time.Now().Unix() > int64(claims["exp"].(float64)) {
				httputil.Unauthorized[string](c, "Token expired")
				return
			}
			id, ok := claims["sub"]
			if !ok {
				httputil.Unauthorized[string](c, "Invalid token")
				return
			}
			user, err := repositories.UserRepo().GetUser(int(id.(float64)))
			if err != nil {
				httputil.Unauthorized[string](c, "Invalid token")
				return
			}
			c.Set("user", user)
			c.Next()
		} else {
			httputil.Unauthorized[string](c, "Invalid token")
			return
		}
		c.Next()
	}
}

func Admin() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("user").(*models.User)
		if user.Username != os.Getenv("ADMIN_USERNAME") {
			httputil.Forbidden[string](c, "You are not allowed to access this resource")
			return
		}
		c.Next()
	}
}
