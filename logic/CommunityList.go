package logic

import (
	"practice/dao/mysql"
	"practice/models"
)

// CommunityList 社区列表查询
func CommunityList() ([]*models.Communities, error) {
	list, err := mysql.CommunityList()
	if err != nil {
		return nil, err
	}
	var communityList []*models.Communities
	for _, c := range list {
		communityList = append(communityList, &c)
	}
	return communityList, nil

}

func ArticleSearch(categoryName string) ([]models.ParamArticle, error) {
	//文章查询
	search, err := mysql.ArticleSearch(categoryName)
	if err != nil {
		return nil, err
	}
	var ArticleSearch []models.ParamArticle
	for _, v := range search {
		ArticleSearch = append(ArticleSearch, v)
	}
	return search, nil
}
