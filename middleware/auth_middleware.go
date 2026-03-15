package middleware

import (
	"net/http"
	"strings"

	"goblog/util"

	"github.com/gin-gonic/gin"
)

func Auth(jwtSecret []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从Header获取Token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			util.Error(c, http.StatusUnauthorized, "Authorization header required")
			c.Abort()
			return
		}

		// 提取Token （Bearer <token>）
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			util.Error(c, http.StatusUnauthorized, "Invalid authorization header format")
			c.Abort()
			return
		}

		tokenString := parts[1]

		//验证Token
		claims, err := util.ParseToken(tokenString, jwtSecret)
		if err != nil {
			util.Error(c, http.StatusUnauthorized, "Invalid token")
			c.Abort()
			return
		}

		//将用户信息存储到Context
		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)

		c.Next()
	}
}
