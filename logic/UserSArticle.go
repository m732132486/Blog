package logic

import (
	"practice/dao/mysql"
	"practice/models"
)

// UsernameSArticle 用户名查询文章
func UsernameSArticle(name string, page, pageSize int) ([]models.UserArticle, error) {
	article, err := mysql.UsernameSArticle(name, page, pageSize)
	if err != nil {
		return nil, err
	}
	return article, nil

}
