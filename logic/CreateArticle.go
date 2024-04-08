package logic

import (
	"practice/dao/mysql"
	"practice/models"
	snowflake "practice/pkg"
)

func CreateArticle(p *models.Article) (err error) {
	id := snowflake.GenID()
	p.TitleID = id
	err = mysql.CreateArticle(p)
	if err != nil {
		return err
	}
	return
}

func Search(keyword string) ([]models.ParamsArticle, error) {
	articles, err := mysql.Search(keyword)
	if err != nil {
		return nil, err
	}
	var ArticleSearch []models.ParamsArticle

	for _, v := range articles {
		ArticleSearch = append(ArticleSearch, v)
	}
	return articles, nil
}

func TitleId(titleID int64) ([]models.ParamsArticle, error) {
	articles, err := mysql.GetArticleByID(titleID)
	if err != nil {
		return nil, err
	}
	var ArticleSearch []models.ParamsArticle

	for _, v := range articles {
		ArticleSearch = append(ArticleSearch, v)
	}
	return articles, nil
}
