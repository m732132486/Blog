package controllers

import (
	"net/http"
	"practice/dao/redis"
	"practice/pkg/jwt"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Like(c *gin.Context) {
	postID := c.Param("id")
	//从token获取用户id
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		zap.L().Error("c.GetHeader(Authorization) ", zap.String("token", tokenString))
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未提供授权令牌"})
		return
	}
	userID, err := jwt.ParseTokenID(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的授权令牌"})
		return
	}
	StrUserId := strconv.FormatInt(userID, 10)
	//检查用户是否点赞过该帖子
	//fmt.Println("postID", reflect.TypeOf(postID))
	//fmt.Println("userid", reflect.TypeOf(StrUserId))
	liked, err := redis.IsLiked(postID, StrUserId)
	if err != nil {
		zap.L().Error("redis.IsLiked(parseInt, userID) ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法检查类似状态"})
		return
	}

	if liked {
		zap.L().Error("redis.IsLiked(parseInt, userID) ", zap.String("postID", postID), zap.String("userID", StrUserId))
		c.JSON(http.StatusBadRequest, gin.H{"error": "你已经点赞了这篇文章"})
		return
	}

	// 执行点赞操作
	err = redis.RecordLike(postID, strconv.FormatInt(userID, 10))
	if err != nil {
		zap.L().Error("redis.RecordLike(parseInt, userID) ", zap.String("parseInt", postID), zap.Int64("userID", userID))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "点赞帖子失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "点赞成功"})
}
