package logic

import (
	"practice/dao/mysql"
	"practice/models"
)

func Delete(Title int64) (err error) {
	return mysql.Delete(Title)

}

func TitleUserid(title int64) ([]models.Article, error) {
	return mysql.TitleUserid(title)

}
