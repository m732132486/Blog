package mysql

import (
	"practice/models"
)

func UsernameSArticle(name string, page, pageSize int) ([]models.UserArticle, error) {
	var article []models.UserArticle
	offset := (page - 1) * pageSize
	if err := DB.Table("articles").
		Select("articles.title,title_id,articles.user_id,articles.parent_id, users.username, communities.category_name,articles.article_content").
		Joins("JOIN users ON articles.user_id = users.user_id").
		Joins("JOIN communities ON articles.parent_id = communities.id").
		Where("users.username LIKE ?", name).
		Offset(offset).
		Limit(pageSize).
		Scan(&article).Error; err != nil {
		return nil, err
	}

	return article, nil

}
