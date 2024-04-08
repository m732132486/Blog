package controllers

import (
	"practice/logic"
	"practice/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func UserFavorites(c *gin.Context) {
	//参数校验
	p := new(models.UserCollect)
	err := c.ShouldBindJSON(p)
	if err != nil {
		zap.L().Error("参数错误", zap.String("err", err.Error()))
		c.JSON(400, gin.H{"code": "参数错误", "msg": err.Error()})
		return
	}
	//业务逻辑
	err = logic.Collect(p)
	if err != nil {
		zap.L().Error("收藏失败", zap.String("err", err.Error()))
		return
	}
	c.JSON(200, gin.H{"code": "ok", "msg": "收藏成功"})

}
func UserFavoritesList(c *gin.Context) {
	//参数校验
	p := new(models.UserCollect)
	err := c.ShouldBindJSON(p)
	if err != nil {
		zap.L().Error("参数错误", zap.String("err", err.Error()))
		c.JSON(400, gin.H{"code": "参数错误", "msg": err.Error()})
		return
	}
	//逻辑处理
	//logic.UserFavoritesList()
	//根据用户id查询收藏列表
	//用户查看收藏列表的内容，获取文章的id进行查看
}
