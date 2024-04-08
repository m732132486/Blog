package controllers

import (
	"net/http"
	"practice/dao/redis"
	"practice/pkg/jwt"
	"strconv"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func CancelLike(c *gin.Context) {
	//获取文章id
	postID := c.Param("id")
	//判断是否有token
	TokenStringId := c.GetHeader("Authorization")
	if TokenStringId == "" {
		zap.L().Error("用户未登录")
		c.JSON(401, gin.H{
			"error": "未提供授权令牌",
		})
		return
	}
	//从token获取用户id
	id, err := jwt.ParseTokenID(TokenStringId)
	if err != nil {
		zap.L().Error("解析token失败", zap.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "登录失败",
		})
		return
	}
	//将用户名转成字符串
	StringId := strconv.FormatInt(id, 10)
	err = redis.CancelLike(postID, StringId)
	if err != nil {
		zap.L().Error("取消点赞失败", zap.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "无法取消",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "取消成功",
	})

}
