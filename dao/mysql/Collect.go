package mysql

import (
	"practice/models"

	"go.uber.org/zap"
)

func Collect(p *models.UserCollect) (err error) {
	if err = DB.Create(p).Error; err != nil {
		zap.L().Error("create user collect error", zap.Any("error", err))
		return err
	}
	return
}
