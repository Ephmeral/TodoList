package middleware

import (
	"github.com/Ephmeral/TodoList/pkg/util"
	"github.com/gin-gonic/gin"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := 200
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404
		} else {
			claim, err := util.ParseToken(token)
			if err != nil {
				code = 403 // token没有权限
			} else if time.Now().Unix() > claim.ExpiresAt {
				code = 401 // token过期无效了
			}
		}
		if code != 200 {
			c.JSON(400, gin.H{
				"status": code,
				"msg":    "Token解析错误",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
