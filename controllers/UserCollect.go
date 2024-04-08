package controllers

import (
	"fmt"
	"net/http"
	"practice/models"
	"practice/pkg/jwt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func UserCollect(c *gin.Context) {
	//参数校验
	p := models.UserCollect{}
	//根据url获取文章id
	//获取token用户id
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		zap.L().Error("c.GetHeader(Authorization) ", zap.String("token", tokenString))
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未提供授权令牌"})
		return
	}
	userID, err := jwt.ParseTokenID(tokenString)
	fmt.Println(userID)
	if err != nil {
		zap.L().Error("GetUser(c) ", zap.Error(err))
		return
	}
	p.UserID = userID
}
