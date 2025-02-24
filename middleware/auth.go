package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rajanlagah/go-course/config"
)



func AuthorizationMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.Abort() // 403
		// read Authorization token from request
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error":"Authorization token not found"})
			c.Abort()
			return
		}
		// remove Bearer from token
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}
		// validate token with salt
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return []byte(config.Config.JWTSaltKey), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error":"Invalid or expired token"})
			c.Abort()
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims) 
		if ok {
			c.Set("email", claims["email"])
			c.Set("name", claims["name"])
		}
		c.Next() // pass
	}
}