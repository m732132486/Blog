package middleware

import (
	"practice/pkg/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuthorization() func(c *gin.Context) {
	return func(c *gin.Context) {
		Authorization := c.Request.Header.Get("Authorization")
		if Authorization == "" {
			c.JSON(401, gin.H{"msg": "请登录"})
			c.Abort()
			return
		}
		// 空格分割
		parts := strings.SplitN(Authorization, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(401, gin.H{"msg": "请登录"})
			c.Abort()
			return
		}
		//parts[1]是获取到tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			c.JSON(401, gin.H{"msg": "请登录"})
			c.Abort()
			return
		}
		//将解析出来的用户信息保存到上下文，后续的中间件和路由就可以使用它了
		c.Set("user", mc.UserID)
		c.Next()

	}
}
