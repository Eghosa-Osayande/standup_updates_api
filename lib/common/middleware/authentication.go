package middleware

import (
	"net/http"
	"os"
	"standup-api/lib/utils/http_response"
	"standup-api/lib/utils/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		splitString := strings.Split(authHeader, "Bearer ")

		if len(splitString) != 2 {

			c.AbortWithStatusJSON(http.StatusUnauthorized, http_response.NewHttpResponseWithError("Unauthorized"))
			return

		}

		token := splitString[1]
		secret, ok := os.LookupEnv("TOKEN_KEY")
		if !ok {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		data, err := jwt.ExtractClaims(token, secret)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, http_response.NewHttpResponseWithError("Unauthorized"))
			return
		}
		role := data.CustomClaims["role"]
		if role != "admin" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, http_response.NewHttpResponseWithError("Unauthorized"))
			return
		}
		c.Next()
	}
}

func EmployeeAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		splitString := strings.Split(authHeader, "Bearer ")

		if len(splitString) != 2 {

			c.AbortWithStatusJSON(http.StatusUnauthorized, http_response.NewHttpResponseWithError("Unauthorized"))
			return

		}

		token := splitString[1]
		secret, ok := os.LookupEnv("TOKEN_KEY")
		if !ok {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		data, err := jwt.ExtractClaims(token, secret)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, http_response.NewHttpResponseWithError("Unauthorized"))
			return
		}
		role := data.CustomClaims["role"]
		if role != "employee" && role != "admin" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, http_response.NewHttpResponseWithError("Unauthorized"))
			return
		}
		c.Set("id", data.CustomClaims["id"])
		c.Next()
	}
}

