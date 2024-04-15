package controllers

import (
	"net/http"
	rediss "practice/dao/redis"
	"practice/logic"
	"practice/models"
	"practice/pkg/jwt"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CommunityHome(c *gin.Context) {
	//查询所有社区

	list, err := logic.CommunityList()
	if err != nil {
		zap.L().Error("查询社区列表失败", zap.Any("err", err))
		c.JSON(500, gin.H{
			"msg": "查询社区列表失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"msg": list,
	})

}

// CommunityList 查询社区文章
func CommunityList(c *gin.Context) {
	// 解析 JSON 请求体
	var CommunitySearchArticles models.CommunitySearchArticles
	if err := c.ShouldBindJSON(&CommunitySearchArticles); err != nil {
		zap.L().Error("解析 JSON 请求体失败", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "解析 JSON 请求体失败"})
		return
	}

	search, err := logic.ArticleSearch(CommunitySearchArticles.CategoryName)

	if err != nil {
		zap.L().Error("查询社区文章失败", zap.Any("err", err))
		return
	}
	//遍历搜索结果
	for i := range search {
		count, err := rediss.GetLikeCount(strconv.FormatInt(search[i].TitleId, 10))
		if err != nil {
			return
		}
		search[i].LikeCount = count
	}
	sort.Slice(search, func(i, j int) bool {
		return search[i].LikeCount > search[j].LikeCount
	})
	c.JSON(http.StatusOK, gin.H{
		"msg": search,
	})
}

// DeleteArticle 删除文章
func DeleteArticle(c *gin.Context) {
	// 解析 JSON 请求体
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("解析请求参数失败", zap.Any("err", err))
		return
	}
	//查找文章的用户id
	titleUserid, err := logic.TitleUserid(id)
	if err != nil {
		return
	}
	//获取用户id
	tokenString := c.GetHeader("Authorization")

	userid, err := jwt.ParseTokenID(tokenString)
	for _, article := range titleUserid {
		if userid != article.UserID {
			zap.L().Error("用户不匹配")
			c.JSON(http.StatusBadRequest, gin.H{
				"err": "用户不匹配",
			})
			return
		}
		err = logic.Delete(id)
		if err != nil {
			zap.L().Error("删除文章失败", zap.Any("err", err))
			return
		}
	}

	c.JSON(200, gin.H{
		"msg": "删除成功",
	})

}
