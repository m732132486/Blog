package mysql

import (
	"fmt"
	"practice/settings"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(conf *settings.MySQLConfig) (err error) {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.DbName)
	dialectic := mysql.Open(dns)
	DB, err = gorm.Open(dialectic, &gorm.Config{})
	if err != nil {
		zap.L().Error("mysql init failed", zap.Any("err", err))
		return err
	}
	DB, err := DB.DB()
	if err != nil {
		zap.L().Error("mysql init failed", zap.Any("err", err))
		return
	}
	DB.SetMaxIdleConns(conf.MaxIdleConns)
	DB.SetMaxOpenConns(conf.MaxOpenConns)
	return nil
}
