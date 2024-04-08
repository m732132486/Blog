package controllers

import (
	"practice/dao/redis"
	"practice/logic"
	"practice/models"
	"sort"
	"strconv"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// UserSArticle 用户创建的文章
func UserSArticle(c *gin.Context) {
	//参数验证
	var name models.UsernameSArticle
	err := c.ShouldBindJSON(&name)
	if err != nil {
		zap.L().Error("参数验证失败", zap.Any("err", err))
		c.JSON(401, gin.H{"msg": "参数验证失败"})
		return
	}
	//查询自己的文章id查询已创建的文章
	article, err := logic.UsernameSArticle(name.Username, 1, 5)
	if err != nil {
		return
	}
	for i := range article {
		count, err := redis.GetLikeCount(strconv.FormatInt(article[i].TitleID, 10))
		if err != nil {
			return
		}
		article[i].LikeCount = count
	}
	sort.Slice(article, func(i, j int) bool {
		return article[i].LikeCount > article[j].LikeCount
	})

	c.JSON(200, gin.H{"article": article, "page": name.Page})

}
