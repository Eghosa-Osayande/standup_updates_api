package middleware

import (
	"errors"
	"net/http"
	"os"
	"standup-api/lib/common/dto"
	"standup-api/lib/utils/http_response"
	"standup-api/lib/utils/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func extractClaims(authHeader string) (*jwt.JWTData, error) {

	splitString := strings.Split(authHeader, "Bearer ")

	if len(splitString) != 2 {
		return nil, errors.New("invalid token")
	}

	token := splitString[1]
	secret, ok := os.LookupEnv("TOKEN_KEY")
	if !ok {
		return nil, errors.New("something went wrong")
	}

	data, err := jwt.ExtractClaims(token, secret)
	return data, err
}

func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		data, err := extractClaims(authHeader)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, http_response.NewHttpResponseWithError(err.Error()))
			return
		}

		role := data.CustomClaims["role"].(string)
		if role != dto.RoleAdmin {
			c.AbortWithStatusJSON(http.StatusUnauthorized, http_response.NewHttpResponseWithError("Unauthorized"))
			return
		}

		c.Set("user_id", data.CustomClaims["id"])
		c.Next()
	}
}

func EmployeeAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		data, err := extractClaims(authHeader)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, http_response.NewHttpResponseWithError(err.Error()))
			return
		}

		role := data.CustomClaims["role"].(string)
		if role != dto.RoleAdmin && role != dto.RoleEmployee {
			c.AbortWithStatusJSON(http.StatusUnauthorized, http_response.NewHttpResponseWithError("Unauthorized"))
			return
		}
		c.Set("user_id", data.CustomClaims["id"])
		c.Next()
	}
}
