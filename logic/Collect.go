package logic

import (
	"practice/dao/mysql"
	"practice/models"
)

func Collect(p *models.UserCollect) (err error) {
	err = mysql.Collect(p)
	if err != nil {
		return err
	}
	return nil
}

/*func UserFavoritesList(p *models.UserCollect) ([]models.UserCollect, error) {
	mysql.UserFavoritesList(p)
}*/
