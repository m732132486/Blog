package logic

import "practice/dao/mysql"

func Delete(Title int64) (err error) {
	return mysql.Delete(Title)

}
