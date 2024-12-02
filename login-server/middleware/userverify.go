package middleware

import (
	"login-server/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func UserVerifyMiddleware(jwtUtil *utils.JWTUtil) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 Authorization 头

		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		// 分离 Bearer 和 token
		parts := strings.SplitN(tokenString, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"})
			c.Abort()
			return
		}

		// 解析 JWT
		id, err := jwtUtil.ParseToken(parts[1]) // parts[1] 是实际的 token
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// 将解析后的用户 ID 存储到上下文中
		c.Set("userID", id)

		c.Next() // 继续处理请求
	}
}
