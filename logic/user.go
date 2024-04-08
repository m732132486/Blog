package logic

import (
	"go.uber.org/zap"
	"practice/dao/mysql"
	"practice/models"
	snowflake "practice/pkg"
	"practice/pkg/jwt"
)

func SignUp(p *models.Params) (err error) {
	//1.判断用户名是否已经存在
	err = mysql.CheckIfTheUserExists(p.Username)
	if err != nil {
		//用户名已经存在
		zap.L().Error("mysql.CheckIfTheUserExists", zap.Error(err))
		return
	}
	//2。生成uid
	ID := snowflake.GenID()
	//构造一个user
	user := &models.User{
		UserID:   ID,
		Username: p.Username,
		Password: p.Password,
	}
	//3.密码加密
	//4.保存进数据库
	err = mysql.InsertUser(user)
	if err != nil {
		zap.L().Error("mysql.InsertUser(user)", zap.Error(err))
		return
	}
	return nil
}

func Login(p *models.ParamsLogin) (user *models.User, err error) {
	user = &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	err = mysql.Login(user)
	if err != nil {
		return nil, err
	}
	//生成token
	generateJWT, err := jwt.GenerateJWT(user.Username, user.UserID)
	if err != nil {
		zap.L().Error("jwt.GenerateJWT", zap.Error(err))
		return nil, err
	}
	user.Token = generateJWT
	return
}
