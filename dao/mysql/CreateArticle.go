package mysql

import (
	"practice/models"

	"go.uber.org/zap"
)

func CreateArticle(p *models.Article) (err error) {
	err = DB.Create(&p).Error
	if err != nil {
		zap.L().Error("create article error", zap.Any("err", err))
		return
	}
	return nil
}
func Search(keyword string) ([]models.ParamsArticle, error) {
	var articles []models.ParamsArticle
	if err := DB.Table("articles").
		Select("articles.title_id, articles.title, articles.user_id, articles.parent_id, articles.article_content, articles.created_at, articles.updated_at, users.username, communities.category_name").
		Joins("JOIN users ON articles.user_id = users.user_id").
		Joins("JOIN communities ON articles.parent_id = communities.id").
		Where("articles.title LIKE ?", "%"+keyword+"%").
		Scan(&articles).Error; err != nil {
	}
	return articles, nil

}

// GetArticleByID 根据文章ID查找文章
func GetArticleByID(titleID int64) ([]models.ParamsArticle, error) {
	var articles []models.ParamsArticle
	if err := DB.Table("articles").
		Select("articles.title_id, articles.title, articles.user_id, articles.parent_id, articles.article_content, articles.created_at, articles.updated_at, users.username, communities.category_name").
		Joins("JOIN users ON articles.user_id = users.user_id").
		Joins("JOIN communities ON articles.parent_id = communities.id").
		Where("articles.title_id = ?", titleID).
		Scan(&articles).Error; err != nil {
	}
	return articles, nil

}
