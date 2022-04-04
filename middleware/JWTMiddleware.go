package middleware

import (
	"fmt"
	"net/http"
	"pustaka-api/book"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenString string
		authHeader := strings.Split(c.GetHeader("Authorization"), " ")
		if len(authHeader) == 2 {
			tokenString = authHeader[1]
		}

		if string(authHeader[0]) != "Bearer" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token, err := book.JWTAuthService().ValidateToken(tokenString)

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			claims.Valid()
		} else {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}

	}

}
