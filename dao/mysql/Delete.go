package mysql

import (
	"errors"
	"practice/models"

	"gorm.io/gorm"
)

func Delete(TitleID int64) error {
	var article models.Article
	result := DB.First(&article, "title_id = ?", TitleID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("文章不存在")
		}
		return result.Error
	}

	result = DB.Delete(&article)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func TitleUserid(titleID int64) ([]models.Article, error) {
	var articles []models.Article
	if err := DB.Where("title_id = ?", titleID).Find(&articles).Error; err != nil {
		return nil, err
	}

	return articles, nil
}
