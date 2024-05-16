package middleware

import (
	service "DailyEnglish/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func tokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		fmt.Println(authHeader)
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, "未提供令牌")
			c.Abort()
			return
		}
		// 检查头是否以"Bearer"开头
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, "令牌格式错误")
			c.Abort()
			return
		}
		// 提取令牌
		token := strings.TrimPrefix(authHeader, "Bearer ")
		user, err := service.ParseToken(token)
		fmt.Println(user, token)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusUnauthorized, "令牌无效")
			c.Abort()
			return
		}

		// 将用户信息存储在context中，后续的处理器可以使用
		c.Set("user", user)
		c.Next()
	}
}
