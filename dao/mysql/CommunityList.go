package mysql

import (
	"practice/models"
)

// CommunityList 社区列表查询
func CommunityList() ([]models.Communities, error) {
	var list []models.Communities
	if err := DB.Select([]string{"category_id", "category_name"}).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func ArticleSearch(categoryName string) ([]models.ParamArticle, error) {
	var articles []models.ParamArticle
	if err := DB.Table("articles AS a").
		Select("a.title AS article_title, a.title_id, u.username AS username, u.user_id, a.parent_id AS parent_id,"+
			" a.article_content AS article_content, a.created_at, a.updated_at, "+
			"c.category_name AS category_name").
		Joins("JOIN users AS u ON a.user_id = u.user_id").
		Joins("JOIN communities AS c ON a.parent_id = c.id").
		Where("c.category_name = ?", categoryName).
		Find(&articles).Error; err != nil {
		// 查询出错
		return nil, err
	}
	return articles, nil
}
