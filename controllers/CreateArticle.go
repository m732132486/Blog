package controllers

import (
	"net/http"
	"practice/dao/redis"
	"practice/logic"
	"practice/models"
	"practice/pkg/jwt"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateArticle(c *gin.Context) {
	// 1. 参数校验
	p := new(models.Article)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("c.ShouldBindJSON(p) ", zap.Error(err))
		c.JSON(200, gin.H{"code": 401, "msg": "参数错误"})
		return
	}
	// 2. 业务处理
	//获取用户id
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		zap.L().Error("c.GetHeader(Authorization) ", zap.String("token", tokenString))
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未提供授权令牌"})
		return
	}
	userID, err := jwt.ParseTokenID(tokenString)
	if err != nil {
		zap.L().Error("GetUser(c) ", zap.Error(err))
		return
	}
	p.UserID = userID
	err = logic.CreateArticle(p)
	if err != nil {
		return
	}
	c.JSON(200, gin.H{"code": "创建成功"})
}

func ArticleSearch(c *gin.Context) {
	var keyword models.Query
	if err := c.ShouldBindJSON(&keyword); err != nil {
		zap.L().Error("c.ShouldBindJSON(&keyword) ", zap.Error(err))
		c.JSON(200, gin.H{"code": 401, "msg": "参数错误"})
		return
	}
	articles, err := logic.Search(keyword.Keyword)
	if err != nil {
		zap.L().Error("logic.Search(keyword.Keyword) ", zap.Error(err))
		return
	}
	for i := range articles {
		likeCount, err := redis.GetLikeCount(strconv.FormatInt(articles[i].TitleID, 10))
		if err != nil {
			return
		}
		articles[i].LikeCount = likeCount
	}
	sort.Slice(articles, func(i, j int) bool {
		return articles[i].LikeCount > articles[j].LikeCount
	})
	c.JSON(200, gin.H{"code": 200, "data": articles})
}

func TitleId(c *gin.Context) {
	//获取文章id
	titleID := c.Param("id")
	//根据id文章搜索
	i, err := strconv.ParseInt(titleID, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"code": 400, "message": "无效的ID格式"})
		return
	}
	val, err := logic.TitleId(i)
	if err != nil {
		return
	}
	//获取点赞计数
	likeCount, err := redis.GetLikeCount(titleID)
	if err != nil {
		zap.L().Error("获取点赞计数失败", zap.Error(err))
		c.JSON(200, gin.H{"code": 401, "msg": "参数错误"})
		return
	}
	val[0].LikeCount = likeCount
	c.JSON(200, gin.H{"code": 200, "data": val})
	//返回文章详情
}
