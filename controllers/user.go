package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"practice/logic"
	"practice/models"
)

func SignUpHandler(c *gin.Context) {
	//1.获取参数校验
	p := new(models.Params)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("c.ShouldBindJSON", zap.Error(err))
		c.JSON(200, gin.H{"code": 1, "msg": "参数错误"})
		return
	}
	//2.业务处理
	err := logic.SignUp(p)
	if err != nil {
		zap.L().Error("logic.SignUp", zap.Error(err))
		c.JSON(200, gin.H{"code": 1, "msg": "注册失败"})
		return
	}
	//3.返回结果
	c.JSON(200, gin.H{"code": 0, "msg": "注册成功"})
}
func LoginHandler(c *gin.Context) {
	//1 登录请求参数校验
	p := new(models.ParamsLogin)
	err := c.ShouldBindJSON(p)
	if err != nil {
		zap.L().Error("c.ShouldBindJSON", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": "参数错误"})
		return
	}
	//2.业务处理
	login, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic.Login", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": "登录失败"})
		return
	}

	//3.返回响应
	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"msg":   "登录成功",
		"data":  login.Username,
		"token": login.Token,
	})

}
