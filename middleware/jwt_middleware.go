package middleware

import (
	"loom/helper"
	"loom/model/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := helper.ExtractTokenFromHeader(c)

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, response.BaseResponseDTO{
				Message:    "Authorization token not provided",
				StatusCode: http.StatusUnauthorized,
			})
			c.Abort()
			return
		}

		claims := &helper.JWTClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return helper.SecretKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, response.BaseResponseDTO{
				Message:    "Invalid token",
				StatusCode: http.StatusUnauthorized,
			})
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
