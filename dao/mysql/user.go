package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"practice/models"
)

func CheckIfTheUserExists(username string) (err error) {
	//检查指定用户是否存在
	err = DB.Where("username = ?", username).Error
	if err == nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			zap.L().Error("CheckIfTheUserExists", zap.String("username", username), zap.Error(err))
			return err
		}

	}
	return nil
}

// encryptPassword 密码加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(oPassword))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
func InsertUser(user *models.User) (err error) {
	//密码加密
	user.Password = encryptPassword(user.Password)
	//插入用户
	if err := DB.Create(&user).Error; err != nil {
		zap.L().Error("DB.Create", zap.Error(err))
		return err
	}
	return

}
func Login(user *models.User) (err error) {
	oPassword := user.Password
	//1.将用户密码转成密文
	password := encryptPassword(oPassword)
	//2.查询数据库
	err = DB.Where("username=? and password=?", user.Username, password).First(&user).Error
	if err != nil {
		zap.L().Error("用户或密码错误", zap.Error(err))
		return err
	}
	return
}
